package maxsubarray

import (
	"reflect"
	"testing"
)

func TestFindMaxSubarray(t *testing.T) {
	var tests = []struct {
		name       string
		arr        []int
		wantSubarr []int
		wantSum    int
	}{
		{
			name:       "mixed ints",
			arr:        []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			wantSubarr: []int{4, -1, 2, 1},
			wantSum:    6,
		},
		{
			name:       "all positive ints",
			arr:        []int{2, 1, 3, 8, 4},
			wantSubarr: []int{2, 1, 3, 8, 4},
			wantSum:    18,
		},
		{
			name:       "all negative ints",
			arr:        []int{-2, -1, -3, -8, -4},
			wantSubarr: []int{},
			wantSum:    0,
		},
		{
			name:       "empty array",
			arr:        []int{},
			wantSubarr: []int{},
			wantSum:    0,
		},
		{
			name:       "Dmitry's acceptance test",
			arr:        []int{-2, 1, -3, 4, -1, 2, 1, -15, 4, 1, -9},
			wantSubarr: []int{4, -1, 2, 1},
			wantSum:    6,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			subarr, sum := FindMaxSubarray(tc.arr)
			if sum != tc.wantSum || !reflect.DeepEqual(subarr, tc.wantSubarr) {
				t.Errorf("Want %v = %v, got %v = %v",
					tc.wantSubarr, tc.wantSum, subarr, sum)
			}
		})
	}
}
