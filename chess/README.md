# Chess problem: count safe boards

Count all distinct placements of chess pieces on an M×N board where none of the pieces is in a position to capture any other piece. Assume the color of the piece does not matter.

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

#### Acceptance test

- Input: 6×9 board with 2 Kings, 1 Queen, 1 Bishop, 1 Rook and 1 Knight
- Result: 20,136,752

### Benchmark

I benchmarked and optimized the code running acceptance test on my MBP using `go test`, `pprof` and _Activity Monitor_:

| Version                       | Time   | Max Mem | Total Mem | Allocs  |
| :---------------------------- | -----: | ------: | --------: | ------: |
| Sequential brute force        | 3m 24s | 48.0 MB |   64.0 GB | 1,995 M |
| Use `int8` wherever possible  | 2m 50s | 33.0 MB |   40.6 GB | 1,995 M |
| Allocate squares in 2 allocs  | 2m 18s | 31.3 MB |   37.4 GB | 1,408 M |
| Use 2 bits per square         | 1m 43s | 19.4 MB |   16.6 GB | 1,291 M |
| Use `[5]int8` to store pieces | 1m 05s | 14.0 MB |   11.6 GB | 1,151 M |
| Optimize `PlacePiece()`       |    49s | 12.5 MB |    9.4 GB |   469 M |
| Use Board Pool                |    33s |  9.0 MB |    2.0 GB |   235 M |

#### Cheat sheet

```
go generate ./...
go test ./...
go test -run=XX -bench=BenchmarkCountSafeBoardsHeavy -benchmem \
    -cpuprofile=cpu.prof -memprofile=mem.prof
go tool pprof chess.test cpu.prof
go tool pprof --alloc_space chess.test mem.prof
```
