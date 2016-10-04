package popcount

import (
	"math/rand"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for _, algo := range algos {
		b.Run(algo.name, func(b *testing.B) {
			benchmarkCount(b, algo.count)
		})
	}
}

func benchmarkCount(b *testing.B, count CountFunc) {
	values := randomValues(1e5)
	b.ResetTimer()
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
