package popcount

import "testing"

var algos = []struct {
	name  string
	count CountFunc
}{
	{"TotallyNaive", CountTotallyNaive},
	{"Naive", CountNaive},
	{"Kernighan", CountKernighan},
	{"MapLookup8", CountMapLookup8},
	{"MapLookup16", CountMapLookup16},
	{"TableLookup8", CountTableLookup8},
	{"TableLookup16", CountTableLookup16},
	{"ParallelNaive", CountParallelNaive},
	{"ParallelSmart", CountParallelSmart},
	{"ParallelSmartNoMul", CountParallelSmartNoMul},
}

func TestCount(t *testing.T) {
	for _, algo := range algos {
		t.Run(algo.name, func(t *testing.T) {
			testCount(t, algo.count)
		})
	}
}

func testCount(t *testing.T, count CountFunc) {
	tests := []struct {
		x, c uint32
	}{
		{0x00000000, 0},
		{0x02468ace, 12},
		{0x13579bdf, 20},
	}
	for _, tc := range tests {
		c := count(tc.x)
		assertCount(t, c, tc.c)
	}
}

func TestCountSlice(t *testing.T) {
	fakeCount := func(x uint32) uint32 { return x }
	c := CountSlice([]uint32{1, 3, 5, 8, 13}, fakeCount)
	assertCount(t, c, 30)
}

func assertCount(t *testing.T, got, want uint32) {
	if got != want {
		t.Errorf("Want %d, got %d", want, got)
	}
}
