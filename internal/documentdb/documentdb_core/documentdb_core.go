// Code generated by "genwrap -debug -schemas=documentdb_api,documentdb_api_catalog,documentdb_api_internal,documentdb_core"; DO NOT EDIT.

package documentdb_core

import (
	"context"
	"log/slog"

	"github.com/FerretDB/wire/wirebson"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/FerretDB/FerretDB/v2/internal/mongoerrors"
)

// BsonCompare is a wrapper for
//
//	documentdb_core.bson_compare(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, OUT bson_compare integer).
func BsonCompare(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument) (outBsonCompare int32, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_compare", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_compare FROM documentdb_core.bson_compare($1::bytea, $2::bytea)", anonymous, anonymous1)
	if err = row.Scan(&outBsonCompare); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_compare", l)
	}
	return
}

// BsonEqual is a wrapper for
//
//	documentdb_core.bson_equal(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, OUT bson_equal boolean).
func BsonEqual(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument) (outBsonEqual bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_equal", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_equal FROM documentdb_core.bson_equal($1::bytea, $2::bytea)", anonymous, anonymous1)
	if err = row.Scan(&outBsonEqual); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_equal", l)
	}
	return
}

// BsonFromBytea is a wrapper for
//
//	documentdb_core.bson_from_bytea(anonymous bytea, OUT bson_from_bytea documentdb_core.bson).
func BsonFromBytea(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonFromBytea wirebson.RawDocument, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_from_bytea", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_from_bytea::bytea FROM documentdb_core.bson_from_bytea($1)", anonymous)
	if err = row.Scan(&outBsonFromBytea); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_from_bytea", l)
	}
	return
}

// BsonGetValue is a wrapper for
//
//	documentdb_core.bson_get_value(anonymous documentdb_core.bson, anonymous1 text, OUT bson_get_value documentdb_core.bson).
func BsonGetValue(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 string) (outBsonGetValue wirebson.RawDocument, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_get_value", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_get_value::bytea FROM documentdb_core.bson_get_value($1::bytea, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonGetValue); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_get_value", l)
	}
	return
}

// BsonGetValueText is a wrapper for
//
//	documentdb_core.bson_get_value_text(anonymous documentdb_core.bson, anonymous1 text, OUT bson_get_value_text text).
func BsonGetValueText(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 string) (outBsonGetValueText string, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_get_value_text", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_get_value_text FROM documentdb_core.bson_get_value_text($1::bytea, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonGetValueText); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_get_value_text", l)
	}
	return
}

// BsonGt is a wrapper for
//
//	documentdb_core.bson_gt(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, OUT bson_gt boolean).
func BsonGt(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument) (outBsonGt bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_gt", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_gt FROM documentdb_core.bson_gt($1::bytea, $2::bytea)", anonymous, anonymous1)
	if err = row.Scan(&outBsonGt); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_gt", l)
	}
	return
}

// BsonGte is a wrapper for
//
//	documentdb_core.bson_gte(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, OUT bson_gte boolean).
func BsonGte(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument) (outBsonGte bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_gte", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_gte FROM documentdb_core.bson_gte($1::bytea, $2::bytea)", anonymous, anonymous1)
	if err = row.Scan(&outBsonGte); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_gte", l)
	}
	return
}

// BsonHashInt4 is a wrapper for
//
//	documentdb_core.bson_hash_int4(anonymous documentdb_core.bson, OUT bson_hash_int4 integer).
func BsonHashInt4(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument) (outBsonHashInt4 int32, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_hash_int4", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_hash_int4 FROM documentdb_core.bson_hash_int4($1::bytea)", anonymous)
	if err = row.Scan(&outBsonHashInt4); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_hash_int4", l)
	}
	return
}

// BsonHashInt8 is a wrapper for
//
//	documentdb_core.bson_hash_int8(anonymous documentdb_core.bson, anonymous1 bigint, OUT bson_hash_int8 bigint).
func BsonHashInt8(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 int64) (outBsonHashInt8 int64, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_hash_int8", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_hash_int8 FROM documentdb_core.bson_hash_int8($1::bytea, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonHashInt8); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_hash_int8", l)
	}
	return
}

// BsonHexToBson is a wrapper for
//
//	documentdb_core.bson_hex_to_bson(anonymous cstring, OUT bson_hex_to_bson documentdb_core.bson).
func BsonHexToBson(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonHexToBson wirebson.RawDocument, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_hex_to_bson", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_hex_to_bson::bytea FROM documentdb_core.bson_hex_to_bson($1)", anonymous)
	if err = row.Scan(&outBsonHexToBson); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_hex_to_bson", l)
	}
	return
}

// BsonIn is a wrapper for
//
//	documentdb_core.bson_in(anonymous cstring, OUT bson_in documentdb_core.bson).
func BsonIn(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonIn wirebson.RawDocument, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_in", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_in::bytea FROM documentdb_core.bson_in($1)", anonymous)
	if err = row.Scan(&outBsonIn); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_in", l)
	}
	return
}

// BsonInRangeInterval is a wrapper for
//
//	documentdb_core.bson_in_range_interval(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, anonymous12 interval, anonymous123 boolean, anonymous1234 boolean, OUT bson_in_range_interval boolean).
func BsonInRangeInterval(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument, anonymous12 struct{}, anonymous123 bool, anonymous1234 bool) (outBsonInRangeInterval bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_in_range_interval", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_in_range_interval FROM documentdb_core.bson_in_range_interval($1::bytea, $2::bytea, $3, $4, $5)", anonymous, anonymous1, anonymous12, anonymous123, anonymous1234)
	if err = row.Scan(&outBsonInRangeInterval); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_in_range_interval", l)
	}
	return
}

// BsonInRangeNumeric is a wrapper for
//
//	documentdb_core.bson_in_range_numeric(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, anonymous12 documentdb_core.bson, anonymous123 boolean, anonymous1234 boolean, OUT bson_in_range_numeric boolean).
func BsonInRangeNumeric(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument, anonymous12 wirebson.RawDocument, anonymous123 bool, anonymous1234 bool) (outBsonInRangeNumeric bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_in_range_numeric", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_in_range_numeric FROM documentdb_core.bson_in_range_numeric($1::bytea, $2::bytea, $3::bytea, $4, $5)", anonymous, anonymous1, anonymous12, anonymous123, anonymous1234)
	if err = row.Scan(&outBsonInRangeNumeric); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_in_range_numeric", l)
	}
	return
}

// BsonJsonToBson is a wrapper for
//
//	documentdb_core.bson_json_to_bson(anonymous text, OUT bson_json_to_bson documentdb_core.bson).
func BsonJsonToBson(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous string) (outBsonJsonToBson wirebson.RawDocument, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_json_to_bson", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_json_to_bson::bytea FROM documentdb_core.bson_json_to_bson($1)", anonymous)
	if err = row.Scan(&outBsonJsonToBson); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_json_to_bson", l)
	}
	return
}

// BsonLt is a wrapper for
//
//	documentdb_core.bson_lt(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, OUT bson_lt boolean).
func BsonLt(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument) (outBsonLt bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_lt", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_lt FROM documentdb_core.bson_lt($1::bytea, $2::bytea)", anonymous, anonymous1)
	if err = row.Scan(&outBsonLt); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_lt", l)
	}
	return
}

// BsonLte is a wrapper for
//
//	documentdb_core.bson_lte(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, OUT bson_lte boolean).
func BsonLte(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument) (outBsonLte bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_lte", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_lte FROM documentdb_core.bson_lte($1::bytea, $2::bytea)", anonymous, anonymous1)
	if err = row.Scan(&outBsonLte); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_lte", l)
	}
	return
}

// BsonNotEqual is a wrapper for
//
//	documentdb_core.bson_not_equal(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, OUT bson_not_equal boolean).
func BsonNotEqual(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument) (outBsonNotEqual bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_not_equal", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_not_equal FROM documentdb_core.bson_not_equal($1::bytea, $2::bytea)", anonymous, anonymous1)
	if err = row.Scan(&outBsonNotEqual); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_not_equal", l)
	}
	return
}

// BsonObjectKeys is a wrapper for
//
//	documentdb_core.bson_object_keys(anonymous documentdb_core.bson, OUT bson_object_keys text).
func BsonObjectKeys(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument) (outBsonObjectKeys string, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_object_keys", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_object_keys FROM documentdb_core.bson_object_keys($1::bytea)", anonymous)
	if err = row.Scan(&outBsonObjectKeys); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_object_keys", l)
	}
	return
}

// BsonOperatorSelectivity is a wrapper for
//
//	documentdb_core.bson_operator_selectivity(anonymous internal, anonymous1 oid, anonymous12 internal, anonymous123 integer, OUT bson_operator_selectivity double precision).
func BsonOperatorSelectivity(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}, anonymous1 struct{}, anonymous12 struct{}, anonymous123 int32) (outBsonOperatorSelectivity float64, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_operator_selectivity", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_operator_selectivity FROM documentdb_core.bson_operator_selectivity($1, $2, $3, $4)", anonymous, anonymous1, anonymous12, anonymous123)
	if err = row.Scan(&outBsonOperatorSelectivity); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_operator_selectivity", l)
	}
	return
}

// BsonOut is a wrapper for
//
//	documentdb_core.bson_out(anonymous documentdb_core.bson, OUT bson_out cstring).
func BsonOut(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument) (outBsonOut struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_out", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_out FROM documentdb_core.bson_out($1::bytea)", anonymous)
	if err = row.Scan(&outBsonOut); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_out", l)
	}
	return
}

// BsonRecv is a wrapper for
//
//	documentdb_core.bson_recv(anonymous internal, OUT bson_recv documentdb_core.bson).
func BsonRecv(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonRecv wirebson.RawDocument, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_recv", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_recv::bytea FROM documentdb_core.bson_recv($1)", anonymous)
	if err = row.Scan(&outBsonRecv); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_recv", l)
	}
	return
}

// BsonRepathAndBuild is a wrapper for
//
//	documentdb_core.bson_repath_and_build(anonymous "any", OUT bson_repath_and_build documentdb_core.bson).
func BsonRepathAndBuild(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonRepathAndBuild wirebson.RawDocument, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_repath_and_build", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_repath_and_build::bytea FROM documentdb_core.bson_repath_and_build($1)", anonymous)
	if err = row.Scan(&outBsonRepathAndBuild); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_repath_and_build", l)
	}
	return
}

// BsonSend is a wrapper for
//
//	documentdb_core.bson_send(anonymous documentdb_core.bson, OUT bson_send bytea).
func BsonSend(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument) (outBsonSend struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_send", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_send FROM documentdb_core.bson_send($1::bytea)", anonymous)
	if err = row.Scan(&outBsonSend); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_send", l)
	}
	return
}

// BsonToBsonHex is a wrapper for
//
//	documentdb_core.bson_to_bson_hex(anonymous documentdb_core.bson, OUT bson_to_bson_hex cstring).
func BsonToBsonHex(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument) (outBsonToBsonHex struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_to_bson_hex", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_to_bson_hex FROM documentdb_core.bson_to_bson_hex($1::bytea)", anonymous)
	if err = row.Scan(&outBsonToBsonHex); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_to_bson_hex", l)
	}
	return
}

// BsonToBsonsequence is a wrapper for
//
//	documentdb_core.bson_to_bsonsequence(anonymous documentdb_core.bson, OUT bson_to_bsonsequence documentdb_core.bsonsequence).
func BsonToBsonsequence(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument) (outBsonToBsonsequence []byte, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_to_bsonsequence", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_to_bsonsequence::bytea FROM documentdb_core.bson_to_bsonsequence($1::bytea)", anonymous)
	if err = row.Scan(&outBsonToBsonsequence); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_to_bsonsequence", l)
	}
	return
}

// BsonToBytea is a wrapper for
//
//	documentdb_core.bson_to_bytea(anonymous documentdb_core.bson, OUT bson_to_bytea bytea).
func BsonToBytea(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument) (outBsonToBytea struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_to_bytea", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_to_bytea FROM documentdb_core.bson_to_bytea($1::bytea)", anonymous)
	if err = row.Scan(&outBsonToBytea); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_to_bytea", l)
	}
	return
}

// BsonToJsonString is a wrapper for
//
//	documentdb_core.bson_to_json_string(anonymous documentdb_core.bson, OUT bson_to_json_string cstring).
func BsonToJsonString(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument) (outBsonToJsonString struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_to_json_string", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_to_json_string FROM documentdb_core.bson_to_json_string($1::bytea)", anonymous)
	if err = row.Scan(&outBsonToJsonString); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_to_json_string", l)
	}
	return
}

// BsonTypanalyze is a wrapper for
//
//	documentdb_core.bson_typanalyze(anonymous internal, OUT bson_typanalyze boolean).
func BsonTypanalyze(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonTypanalyze bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_typanalyze", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_typanalyze FROM documentdb_core.bson_typanalyze($1)", anonymous)
	if err = row.Scan(&outBsonTypanalyze); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_typanalyze", l)
	}
	return
}

// BsonUniqueIndexEqual is a wrapper for
//
//	documentdb_core.bson_unique_index_equal(anonymous documentdb_core.bson, anonymous1 documentdb_core.bson, OUT bson_unique_index_equal boolean).
func BsonUniqueIndexEqual(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 wirebson.RawDocument) (outBsonUniqueIndexEqual bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bson_unique_index_equal", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bson_unique_index_equal FROM documentdb_core.bson_unique_index_equal($1::bytea, $2::bytea)", anonymous, anonymous1)
	if err = row.Scan(&outBsonUniqueIndexEqual); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bson_unique_index_equal", l)
	}
	return
}

// BsonqueryCompare is a wrapper for
//
//	documentdb_core.bsonquery_compare(anonymous documentdb_core.bsonquery, anonymous1 documentdb_core.bsonquery, OUT bsonquery_compare integer).
func BsonqueryCompare(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}, anonymous1 struct{}) (outBsonqueryCompare int32, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_compare", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_compare FROM documentdb_core.bsonquery_compare($1, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonqueryCompare); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_compare", l)
	}
	return
}

// BsonqueryCompare1 is a wrapper for
//
//	documentdb_core.bsonquery_compare(anonymous documentdb_core.bson, anonymous1 documentdb_core.bsonquery, OUT bsonquery_compare integer).
func BsonqueryCompare1(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous wirebson.RawDocument, anonymous1 struct{}) (outBsonqueryCompare int32, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_compare", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_compare FROM documentdb_core.bsonquery_compare($1::bytea, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonqueryCompare); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_compare", l)
	}
	return
}

// BsonqueryEqual is a wrapper for
//
//	documentdb_core.bsonquery_equal(anonymous documentdb_core.bsonquery, anonymous1 documentdb_core.bsonquery, OUT bsonquery_equal boolean).
func BsonqueryEqual(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}, anonymous1 struct{}) (outBsonqueryEqual bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_equal", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_equal FROM documentdb_core.bsonquery_equal($1, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonqueryEqual); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_equal", l)
	}
	return
}

// BsonqueryGt is a wrapper for
//
//	documentdb_core.bsonquery_gt(anonymous documentdb_core.bsonquery, anonymous1 documentdb_core.bsonquery, OUT bsonquery_gt boolean).
func BsonqueryGt(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}, anonymous1 struct{}) (outBsonqueryGt bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_gt", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_gt FROM documentdb_core.bsonquery_gt($1, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonqueryGt); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_gt", l)
	}
	return
}

// BsonqueryGte is a wrapper for
//
//	documentdb_core.bsonquery_gte(anonymous documentdb_core.bsonquery, anonymous1 documentdb_core.bsonquery, OUT bsonquery_gte boolean).
func BsonqueryGte(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}, anonymous1 struct{}) (outBsonqueryGte bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_gte", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_gte FROM documentdb_core.bsonquery_gte($1, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonqueryGte); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_gte", l)
	}
	return
}

// BsonqueryIn is a wrapper for
//
//	documentdb_core.bsonquery_in(anonymous cstring, OUT bsonquery_in documentdb_core.bsonquery).
func BsonqueryIn(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonqueryIn struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_in", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_in FROM documentdb_core.bsonquery_in($1)", anonymous)
	if err = row.Scan(&outBsonqueryIn); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_in", l)
	}
	return
}

// BsonqueryLt is a wrapper for
//
//	documentdb_core.bsonquery_lt(anonymous documentdb_core.bsonquery, anonymous1 documentdb_core.bsonquery, OUT bsonquery_lt boolean).
func BsonqueryLt(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}, anonymous1 struct{}) (outBsonqueryLt bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_lt", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_lt FROM documentdb_core.bsonquery_lt($1, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonqueryLt); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_lt", l)
	}
	return
}

// BsonqueryLte is a wrapper for
//
//	documentdb_core.bsonquery_lte(anonymous documentdb_core.bsonquery, anonymous1 documentdb_core.bsonquery, OUT bsonquery_lte boolean).
func BsonqueryLte(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}, anonymous1 struct{}) (outBsonqueryLte bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_lte", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_lte FROM documentdb_core.bsonquery_lte($1, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonqueryLte); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_lte", l)
	}
	return
}

// BsonqueryNotEqual is a wrapper for
//
//	documentdb_core.bsonquery_not_equal(anonymous documentdb_core.bsonquery, anonymous1 documentdb_core.bsonquery, OUT bsonquery_not_equal boolean).
func BsonqueryNotEqual(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}, anonymous1 struct{}) (outBsonqueryNotEqual bool, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_not_equal", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_not_equal FROM documentdb_core.bsonquery_not_equal($1, $2)", anonymous, anonymous1)
	if err = row.Scan(&outBsonqueryNotEqual); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_not_equal", l)
	}
	return
}

// BsonqueryOut is a wrapper for
//
//	documentdb_core.bsonquery_out(anonymous documentdb_core.bsonquery, OUT bsonquery_out cstring).
func BsonqueryOut(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonqueryOut struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_out", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_out FROM documentdb_core.bsonquery_out($1)", anonymous)
	if err = row.Scan(&outBsonqueryOut); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_out", l)
	}
	return
}

// BsonqueryRecv is a wrapper for
//
//	documentdb_core.bsonquery_recv(anonymous internal, OUT bsonquery_recv documentdb_core.bsonquery).
func BsonqueryRecv(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonqueryRecv struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_recv", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_recv FROM documentdb_core.bsonquery_recv($1)", anonymous)
	if err = row.Scan(&outBsonqueryRecv); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_recv", l)
	}
	return
}

// BsonquerySend is a wrapper for
//
//	documentdb_core.bsonquery_send(anonymous documentdb_core.bsonquery, OUT bsonquery_send bytea).
func BsonquerySend(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonquerySend struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonquery_send", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonquery_send FROM documentdb_core.bsonquery_send($1)", anonymous)
	if err = row.Scan(&outBsonquerySend); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonquery_send", l)
	}
	return
}

// BsonsequenceFromBytea is a wrapper for
//
//	documentdb_core.bsonsequence_from_bytea(anonymous bytea, OUT bsonsequence_from_bytea documentdb_core.bsonsequence).
func BsonsequenceFromBytea(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonsequenceFromBytea []byte, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonsequence_from_bytea", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonsequence_from_bytea::bytea FROM documentdb_core.bsonsequence_from_bytea($1)", anonymous)
	if err = row.Scan(&outBsonsequenceFromBytea); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonsequence_from_bytea", l)
	}
	return
}

// BsonsequenceGetBson is a wrapper for
//
//	documentdb_core.bsonsequence_get_bson(anonymous documentdb_core.bsonsequence, OUT bsonsequence_get_bson documentdb_core.bson).
func BsonsequenceGetBson(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous []byte) (outBsonsequenceGetBson wirebson.RawDocument, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonsequence_get_bson", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonsequence_get_bson::bytea FROM documentdb_core.bsonsequence_get_bson($1::bytea)", anonymous)
	if err = row.Scan(&outBsonsequenceGetBson); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonsequence_get_bson", l)
	}
	return
}

// BsonsequenceIn is a wrapper for
//
//	documentdb_core.bsonsequence_in(anonymous cstring, OUT bsonsequence_in documentdb_core.bsonsequence).
func BsonsequenceIn(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonsequenceIn []byte, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonsequence_in", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonsequence_in::bytea FROM documentdb_core.bsonsequence_in($1)", anonymous)
	if err = row.Scan(&outBsonsequenceIn); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonsequence_in", l)
	}
	return
}

// BsonsequenceOut is a wrapper for
//
//	documentdb_core.bsonsequence_out(anonymous documentdb_core.bsonsequence, OUT bsonsequence_out cstring).
func BsonsequenceOut(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous []byte) (outBsonsequenceOut struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonsequence_out", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonsequence_out FROM documentdb_core.bsonsequence_out($1::bytea)", anonymous)
	if err = row.Scan(&outBsonsequenceOut); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonsequence_out", l)
	}
	return
}

// BsonsequenceRecv is a wrapper for
//
//	documentdb_core.bsonsequence_recv(anonymous internal, OUT bsonsequence_recv documentdb_core.bsonsequence).
func BsonsequenceRecv(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outBsonsequenceRecv []byte, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonsequence_recv", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonsequence_recv::bytea FROM documentdb_core.bsonsequence_recv($1)", anonymous)
	if err = row.Scan(&outBsonsequenceRecv); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonsequence_recv", l)
	}
	return
}

// BsonsequenceSend is a wrapper for
//
//	documentdb_core.bsonsequence_send(anonymous documentdb_core.bsonsequence, OUT bsonsequence_send bytea).
func BsonsequenceSend(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous []byte) (outBsonsequenceSend struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonsequence_send", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonsequence_send FROM documentdb_core.bsonsequence_send($1::bytea)", anonymous)
	if err = row.Scan(&outBsonsequenceSend); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonsequence_send", l)
	}
	return
}

// BsonsequenceToBytea is a wrapper for
//
//	documentdb_core.bsonsequence_to_bytea(anonymous documentdb_core.bsonsequence, OUT bsonsequence_to_bytea bytea).
func BsonsequenceToBytea(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous []byte) (outBsonsequenceToBytea struct{}, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.bsonsequence_to_bytea", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT bsonsequence_to_bytea FROM documentdb_core.bsonsequence_to_bytea($1::bytea)", anonymous)
	if err = row.Scan(&outBsonsequenceToBytea); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.bsonsequence_to_bytea", l)
	}
	return
}

// RowGetBson is a wrapper for
//
//	documentdb_core.row_get_bson(anonymous record, OUT row_get_bson documentdb_core.bson).
func RowGetBson(ctx context.Context, conn *pgx.Conn, l *slog.Logger, anonymous struct{}) (outRowGetBson wirebson.RawDocument, err error) {
	ctx, span := otel.Tracer("").Start(ctx, "documentdb_core.row_get_bson", oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	defer span.End()

	row := conn.QueryRow(ctx, "SELECT row_get_bson::bytea FROM documentdb_core.row_get_bson($1)", anonymous)
	if err = row.Scan(&outRowGetBson); err != nil {
		err = mongoerrors.Make(ctx, err, "documentdb_core.row_get_bson", l)
	}
	return
}
