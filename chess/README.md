# Chess problem: count safe boards

Count all distinct placements of chess pieces on M×N board where none of the pieces is in a position to capture any other piece. Assume the color of the piece does not matter.

- Input:
  - board dimensions
  - number of pieces of each type: King, Queen, Bishop, Rook and Knight
- Output:
  - number of distinct safe boards

### Examples

##### Example 1

- Input: 3×3 board with 2 Kings and 1 Rook
- Result: 4

##### Example 2

- Input: 4×4 board with 2 Rooks and 4 Knights
- Result: 8

### Acceptance test

- Input: 6×9 board with 2 Kings, 1 Queen, 1 Bishop, 1 Rook and 1 Knight
- Result: 20,136,752

### Test

```
go generate ./...
go test ./...
```

### Benchmark acceptance test

I benchmarked and optimized the code on my MBP using `go test` and _Activity Monitor_

```
go test -run XX -bench BenchmarkHeavy -benchmem
```

| Versi                        | Time   | Max Mem | Total Mem | Allocs  |
| :--------------------------- | -----: | ------: | --------: | ------: |
| First version                | 3m 24s | 48.0 MB |   64.0 GB | 1,995 M |
| Use `int8` wherever possible | 2m 50s | 33.0 MB |   40.6 GB | 1,995 M |
