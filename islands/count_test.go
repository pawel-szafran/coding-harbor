package islands

import "testing"

const (
	X = true
	O = false
)

func TestCountIslands(t *testing.T) {
	var tests = []struct {
		name        string
		m           Map
		wantIslands uint
	}{
		{
			name: "no islands",
			m: Map{
				{O, O, O, O, O},
				{O, O, O, O, O},
			},
			wantIslands: 0,
		},
		{
			name: "one-elem island",
			m: Map{
				{O, O, O, O, O},
				{O, O, O, O, O},
				{O, O, X, O, O},
				{O, O, O, O, O},
				{O, O, O, O, O},
			},
			wantIslands: 1,
		},
		{
			name: "one-elem islands",
			m: Map{
				{O, O, O, O, O},
				{O, X, O, O, O},
				{O, O, O, X, O},
			},
			wantIslands: 2,
		},
		{
			name: "many-elems-in-row islands",
			m: Map{
				{X, O, X, X, O},
				{O, O, O, O, X},
				{X, X, X, O, O},
			},
			wantIslands: 4,
		},
		{
			name: "many-elems-in-column islands",
			m: Map{
				{O, O, X, O, O},
				{X, O, X, O, X},
				{O, O, X, O, X},
			},
			wantIslands: 3,
		},
		{
			name: "complex islands",
			m: Map{
				{X, X, O, O, O, O, X, X},
				{O, X, X, O, O, X, X, O},
				{O, X, O, O, O, X, X, O},
				{X, X, X, X, X, O, O, O},
				{O, O, O, O, O, O, O, O},
				{X, O, O, O, O, X, X, O},
				{X, X, O, O, O, X, O, X},
			},
			wantIslands: 5,
		},
		{
			name: "all land",
			m: Map{
				{X, X, X, X},
				{X, X, X, X},
				{X, X, X, X},
				{X, X, X, X},
			},
			wantIslands: 1,
		},
		{
			name: "frame island",
			m: Map{
				{X, X, X, X},
				{X, O, O, X},
				{X, O, O, X},
				{X, X, X, X},
			},
			wantIslands: 1,
		},
		{
			name: "cross island",
			m: Map{
				{O, O, X, O, O},
				{O, X, X, X, O},
				{X, X, X, X, X},
				{O, X, X, X, O},
				{O, O, X, O, O},
			},
			wantIslands: 1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			islands := CountIslands(tc.m)
			if islands != tc.wantIslands {
				t.Errorf("Want %d, got %d", tc.wantIslands, islands)
			}
		})
	}
}
