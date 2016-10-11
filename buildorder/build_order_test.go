package buildorder

import (
	"reflect"
	"testing"
)

func TestFindBuildOrder(t *testing.T) {
	tests := []struct {
		name      string
		modules   []Module
		wantOrder []ModuleName
		wantErr   error
	}{
		{
			name: "single module",
			modules: []Module{
				{Name: "A"},
			},
			wantOrder: []ModuleName{"A"},
		},
		{
			name: "modules with no deps",
			modules: []Module{
				{Name: "A"},
				{Name: "B"},
			},
			wantOrder: []ModuleName{"A", "B"},
		},
		{
			name: "modules with single dep",
			modules: []Module{
				{Name: "A", Deps: []ModuleName{"B"}},
				{Name: "B"},
			},
			wantOrder: []ModuleName{"B", "A"},
		},
		{
			name: "modules with many deps",
			modules: []Module{
				{Name: "A", Deps: []ModuleName{"B", "C"}},
				{Name: "B"},
				{Name: "C"},
			},
			wantOrder: []ModuleName{"B", "C", "A"},
		},
		{
			name: "complex modules graph",
			modules: []Module{
				{Name: "A", Deps: []ModuleName{"B", "C", "F"}},
				{Name: "B", Deps: []ModuleName{"F"}},
				{Name: "C", Deps: []ModuleName{"F"}},
				{Name: "D"},
				{Name: "E", Deps: []ModuleName{"A", "B"}},
				{Name: "F"},
				{Name: "G", Deps: []ModuleName{"D"}},
			},
			wantOrder: []ModuleName{"F", "B", "C", "A", "D", "E", "G"},
		},
		{
			name: "dependency cycle",
			modules: []Module{
				{Name: "A", Deps: []ModuleName{"B"}},
				{Name: "B", Deps: []ModuleName{"C"}},
				{Name: "C", Deps: []ModuleName{"A"}},
			},
			wantErr: ErrDepCycle,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			order, err := FindBuildOrder(tc.modules)
			switch {
			case tc.wantErr != nil && err != tc.wantErr:
				t.Errorf("Want err %q, got %q", tc.wantErr, err)
			case tc.wantErr == nil && err != nil:
				t.Errorf("Unwanted err %q", err)
			case tc.wantOrder != nil && !reflect.DeepEqual(order, tc.wantOrder):
				t.Errorf("Want build order %v, got %v", tc.wantOrder, order)
			}
		})
	}
}
