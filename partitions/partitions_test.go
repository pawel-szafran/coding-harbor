package partitions

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestFindPartitions(t *testing.T) {
	for _, tc := range findPartitionTests {
		t.Run(tc.name, func(t *testing.T) {
			partitions := FindPartitions(tc.graph)
			if !equalPartitions(partitions, tc.partitions) {
				t.Errorf("Want %v, got %v", partitions, tc.partitions)
			}
		})
	}
}

var findPartitionTests = []struct {
	name       string
	graph      Graph
	partitions []Partition
}{
	{
		name: "1-elem partitions",
		graph: Graph{
			Nodes: []Node{1, 2, 3},
			Edges: []Edge{},
		},
		partitions: []Partition{{1}, {2}, {3}},
	},
	{
		name: "2-elem partition",
		graph: Graph{
			Nodes: []Node{1, 2, 3},
			Edges: []Edge{{1, 2}},
		},
		partitions: []Partition{{1, 2}, {3}},
	},
	{
		name: "many-elem partitions",
		graph: Graph{
			Nodes: []Node{1, 2, 3, 4, 5, 6},
			Edges: []Edge{
				{6, 1},
				{2, 3}, {3, 5}, {4, 3},
			},
		},
		partitions: []Partition{{1, 6}, {2, 3, 4, 5}},
	},
	{
		name: "partitions with cycles",
		graph: Graph{
			Nodes: []Node{1, 2, 3, 4, 5, 6, 7},
			Edges: []Edge{
				{5, 1}, {1, 3}, {3, 5},
				{2, 6}, {6, 4}, {4, 7}, {2, 7},
			},
		},
		partitions: []Partition{{5, 1, 3}, {2, 6, 4, 7}},
	},
	{
		name: "1 partition all-to-all",
		graph: Graph{
			Nodes: []Node{1, 2, 3, 4, 5},
			Edges: []Edge{
				{1, 4}, {1, 2}, {1, 5}, {1, 3},
				{4, 2}, {4, 5}, {4, 3},
				{2, 5}, {2, 3},
				{5, 3},
			},
		},
		partitions: []Partition{{1, 2, 3, 4, 5}},
	},
	{
		name:       "empty graph",
		graph:      Graph{},
		partitions: []Partition{},
	},
}

func equalPartitions(a, b []Partition) bool {
	return reflect.DeepEqual(asSet(a), asSet(b))
}

func asSet(partitions []Partition) map[string]bool {
	set := map[string]bool{}
	for _, partition := range partitions {
		set[buildPartitionKey(partition)] = true
	}
	return set
}

func buildPartitionKey(partition Partition) string {
	nodesAsStrings := make([]string, len(partition))
	for i, node := range partition {
		nodesAsStrings[i] = string(node)
	}
	sort.Strings(nodesAsStrings)
	return strings.Join(nodesAsStrings, "-")
}
