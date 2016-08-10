package popcount

import (
	"math/rand"
	"testing"
)

func BenchmarkCountTotallyNaive(b *testing.B)       { benchmarkCount(b, CountTotallyNaive) }
func BenchmarkCountNaive(b *testing.B)              { benchmarkCount(b, CountNaive) }
func BenchmarkCountKernighan(b *testing.B)          { benchmarkCount(b, CountKernighan) }
func BenchmarkCountMapLookup8(b *testing.B)         { benchmarkCount(b, CountMapLookup8) }
func BenchmarkCountMapLookup16(b *testing.B)        { benchmarkCount(b, CountMapLookup16) }
func BenchmarkCountTableLookup8(b *testing.B)       { benchmarkCount(b, CountTableLookup8) }
func BenchmarkCountTableLookup16(b *testing.B)      { benchmarkCount(b, CountTableLookup16) }
func BenchmarkCountParallelNaive(b *testing.B)      { benchmarkCount(b, CountParallelNaive) }
func BenchmarkCountParallelSmart(b *testing.B)      { benchmarkCount(b, CountParallelSmart) }
func BenchmarkCountParallelSmartNoMul(b *testing.B) { benchmarkCount(b, CountParallelSmartNoMul) }

func benchmarkCount(b *testing.B, count CountFunc) {
	values := randomValues(1e5)
	for i := 0; i < b.N; i++ {
		CountSlice(values, count)
	}
}

func randomValues(size int) []uint32 {
	values := make([]uint32, size)
	rand.Seed(8)
	for i := 0; i < size; i++ {
		values[i] = rand.Uint32()
	}
	return values
}
