package missingint

import "testing"

func TestFindMissingInt(t *testing.T) {
	tests := []struct {
		name    string
		ints    []int
		missing int
	}{
		{
			name:    "ordered",
			ints:    []int{1, 2, 3, 5, 6},
			missing: 4,
		},
		{
			name:    "shuffled",
			ints:    []int{3, 6, 1, 5, 2},
			missing: 4,
		},
		{
			name:    "not starting from 1",
			ints:    []int{7, 8, 4, 5},
			missing: 6,
		},
		{
			name:    "large ints",
			ints:    []int{MaxInt - 3, MaxInt, MaxInt - 1},
			missing: MaxInt - 2,
		},
		{
			name:    "negative",
			ints:    []int{-6, -4, -8, -5},
			missing: -7,
		},
		{
			name:    "negative and positive",
			ints:    []int{-3, 0, -1, 2, 1},
			missing: -2,
		},
		{
			name:    "missing 0",
			ints:    []int{-1, 2, -2, 1},
			missing: 0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			missing := FindMissingInt(tc.ints)
			if missing != tc.missing {
				t.Errorf("Missing int in %v is %d, not %d", tc.ints, tc.missing, missing)
			}
		})
	}
}

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)
