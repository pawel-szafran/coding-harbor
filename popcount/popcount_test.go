package popcount

import "testing"

func TestCountTotallyNaive(t *testing.T)  { testCount(t, CountTotallyNaive) }
func TestCountNaive(t *testing.T)         { testCount(t, CountNaive) }
func TestCountKernighan(t *testing.T)     { testCount(t, CountKernighan) }
func TestCountMapLookup8(t *testing.T)    { testCount(t, CountMapLookup8) }
func TestCountMapLookup16(t *testing.T)   { testCount(t, CountMapLookup16) }
func TestCountTableLookup8(t *testing.T)  { testCount(t, CountTableLookup8) }
func TestCountTableLookup16(t *testing.T) { testCount(t, CountTableLookup16) }
func TestCountParallelNaive(t *testing.T) { testCount(t, CountParallelNaive) }
func TestCountParallelSmart(t *testing.T) { testCount(t, CountParallelSmart) }

func testCount(t *testing.T, count CountFunc) {
	tests := []struct {
		x, c uint32
	}{
		{0x00000000, 0},
		{0x02468ace, 12},
		{0x13579bdf, 20},
	}
	for _, tt := range tests {
		c := count(tt.x)
		assertCount(t, c, tt.c)
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
