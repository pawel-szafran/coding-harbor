package popcount

import "testing"

func TestCountNaive(t *testing.T) {
	tests := []struct {
		x, c uint32
	}{
		{0x00000000, 0},
		{0x02468ace, 12},
		{0x13579bdf, 20},
	}
	for _, tt := range tests {
		c := CountNaive(tt.x)
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