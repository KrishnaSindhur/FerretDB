// Copyright 2021 FerretDB Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/FerretDB/FerretDB/v2/internal/util/must"
)

// headerTemplate is used to generate header.
//
// "// Code generated" is intentionally not on the next line
// to prevent generate.go itself being marked as generated on GitHub.
var headerTemplate = template.Must(template.New("header").Parse(`// Code generated by "{{.Cmd}}"; DO NOT EDIT.

package {{.Package}}

import (
	"context"
	"log/slog"

	"github.com/FerretDB/wire/wirebson"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/FerretDB/FerretDB/v2/internal/mongoerrors"
)
`))

// headerData contains information needed for generating header.
type headerData struct {
	Cmd     string
	Package string
}

// funcTemplate is used to generate function.
var funcTemplate = template.Must(template.New("func").Parse(`
// {{.FuncName}} is a wrapper for
//
//	{{.Comment}}.
func {{.FuncName}}({{.Params}}) ({{.Returns}}) {
	ctx, span := otel.Tracer("").Start(ctx, "{{.SQLFuncName}}", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow({{.QueryRowArgs}})
	if err = row.Scan({{.ScanArgs}}); err != nil {
		err = mongoerrors.Make(ctx, err, "{{.SQLFuncName}}", l)
	}
	return
}
`))

// templateData contains information need for generating a function to run SQL query and scan the output.
type templateData struct {
	FuncName     string
	SQLFuncName  string
	Comment      string
	Params       string
	Returns      string
	SQLArgs      string
	SQLReturns   string
	IsProcedure  bool
	QueryRowArgs string
	ScanArgs     string
}

// Generate creates files <schema>/<schema>.go for each schema and generates
// go functions for querying DocumentDB API.
func Generate(schemaRoutines map[string]map[string]templateData) error {
	for schema, routines := range schemaRoutines {
		if err := os.MkdirAll(schema, 0o777); err != nil {
			return err
		}

		out := must.NotFail(os.Create(filepath.Join(schema, schema+".go")))
		defer func() {
			must.NoError(out.Close())
		}()

		h := headerData{
			Cmd:     "genwrap " + strings.Join(os.Args[1:], " "),
			Package: schema,
		}

		if err := headerTemplate.Execute(out, &h); err != nil {
			return err
		}

		for _, k := range slices.Sorted(maps.Keys(routines)) {
			v := routines[k]
			if err := generateGoFunction(out, &v); err != nil {
				return err
			}
		}
	}

	return nil
}

// generateGoFunction uses template data to produce go function for querying DocumentDB API.
//
// The function is generated by using template and written to the writer.
func generateGoFunction(writer io.Writer, data *templateData) error {
	q := generateSQL(data)

	queryRowArgs := fmt.Sprintf(`ctx, "%s"`, q)
	if data.QueryRowArgs != "" {
		queryRowArgs = fmt.Sprintf("%s, %s", queryRowArgs, data.QueryRowArgs)
	}
	data.QueryRowArgs = queryRowArgs

	params := "ctx context.Context, conn *pgx.Conn, l *slog.Logger"
	if data.Params != "" {
		params = fmt.Sprintf("%s, %s", params, data.Params)
	}
	data.Params = params

	returns := "err error"
	if data.Returns != "" {
		returns = fmt.Sprintf("%s, %s", data.Returns, returns)
	}
	data.Returns = returns

	return funcTemplate.Execute(writer, &data)
}

// generateSQL builds SQL query and arguments for the given function definition.
func generateSQL(f *templateData) string {
	if f.IsProcedure {
		return fmt.Sprintf("CALL %s(%s)", f.SQLFuncName, f.SQLArgs)
	}

	if f.SQLReturns == "" {
		return fmt.Sprintf("SELECT FROM %s(%s)", f.SQLFuncName, f.SQLArgs)
	}

	return fmt.Sprintf(
		"SELECT %s FROM %s(%s)",
		f.SQLReturns,
		f.SQLFuncName,
		f.SQLArgs,
	)
}
