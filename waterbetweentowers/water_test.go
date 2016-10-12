package waterbetweentowers

import "testing"

func TestCalcWaterBetweenTowers(t *testing.T) {
	var tests = []struct {
		name      string
		towers    []Unit
		wantWater Unit
	}{
		{
			name:      "simple pool",
			towers:    []Unit{4, 1, 2, 3},
			wantWater: 3,
		},
		{
			name:      "2 pools",
			towers:    []Unit{8, 2, 5, 1, 0, 2},
			wantWater: 6,
		},
		{
			name:      "complex",
			towers:    []Unit{5, 3, 3, 7, 2, 1, 1, 0, 6, 3, 5, 9, 1, 2},
			wantWater: 36,
		},
		{
			name:      "flat",
			towers:    []Unit{2, 2, 2, 2},
			wantWater: 0,
		},
		{
			name:      "empty",
			towers:    []Unit{},
			wantWater: 0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			water := CalcWaterBetweenTowers(tc.towers)
			if water != tc.wantWater {
				t.Errorf("Want %d, got %d", tc.wantWater, water)
			}
		})
	}
}
