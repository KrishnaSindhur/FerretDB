db.runCommand({
  aggregate: 'books',
  pipeline: [
    {
      $search: {
        cosmosSearch: {
          vector: [
            0.02232860028743744, 0.06849973648786545, 0.030828291550278664, 0.0903232991695404, -0.028270352631807327,
            -0.036311957985162735, 0.02430308423936367, -0.051550041884183884, -0.06737732142210007,
            0.011019553989171982, -0.013402754440903664, -0.004793450236320496
          ],
          path: 'vector',
          k: 2,
          efSearch: 40
        }
      }
    },
    {
      $project: {
        title: 1,
        author: {
          $first: '$authors.name'
        },
        summary: 1,
        vector: 1
      }
    }
  ],
  cursor: {}
})
