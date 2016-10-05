# Count set bits (popcount)

Count set bits of 100k 32-bit integers. And benchmark it.

### Algorithms

- Naive
- Brian Kernighan's
- Map lookup
- Table lookup
- Parallel

### Benchmark

```
go test -bench=.
```

Results from my MBP:

```
BenchmarkCount/TotallyNaive-4                100          13679292 ns/op
BenchmarkCount/Naive-4                       500           2635228 ns/op
BenchmarkCount/Kernighan-4                   500           2559293 ns/op
BenchmarkCount/MapLookup8-4                  100          13266373 ns/op
BenchmarkCount/MapLookup16-4                 200           9026950 ns/op
BenchmarkCount/TableLookup8-4               3000            489734 ns/op
BenchmarkCount/TableLookup16-4              3000            456320 ns/op
BenchmarkCount/ParallelNaive-4              3000            558721 ns/op
BenchmarkCount/ParallelSmart-4              3000            455006 ns/op
BenchmarkCount/ParallelSmartNoMul-4         3000            494310 ns/op
PASS
ok      github.com/pawel-szafran/coding-harbor/popcount 16.312s
```
