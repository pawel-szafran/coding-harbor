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
		if c != tt.c {
			t.Errorf("Want %d, got %d", tt.c, c)
		}
	}
}
