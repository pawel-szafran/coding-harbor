package popcount

import (
	"math/rand"
	"testing"
)

func BenchmarkCountTotallyNaive(b *testing.B) { benchmarkCount(b, CountTotallyNaive) }
func BenchmarkCountNaive(b *testing.B)        { benchmarkCount(b, CountNaive) }

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
