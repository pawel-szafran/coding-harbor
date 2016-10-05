package partitions

type Node int
type Edge struct {
	From, To Node
}
type Graph struct {
	Nodes []Node
	Edges []Edge
}
type Partition []Node

func FindPartitions(graph Graph) []Partition {
	return newPartitionFinder(graph).findAll()
}

type partitionFinder struct {
	nodes     []Node
	neighbors map[Node][]Node
	visited   map[Node]bool
}

func newPartitionFinder(graph Graph) *partitionFinder {
	return &partitionFinder{
		nodes:     graph.Nodes,
		neighbors: buildNeighborsMap(graph.Edges),
		visited:   map[Node]bool{},
	}
}

func buildNeighborsMap(edges []Edge) map[Node][]Node {
	neighbors := map[Node][]Node{}
	for _, edge := range edges {
		addNeighbor(neighbors, edge.From, edge.To)
		addNeighbor(neighbors, edge.To, edge.From)
	}
	return neighbors
}

func addNeighbor(neighbors map[Node][]Node, from, to Node) {
	neighbors[from] = append(neighbors[from], to)
}

func (f *partitionFinder) findAll() (partitions []Partition) {
	for _, node := range f.nodes {
		if !f.visited[node] {
			partition := f.visitPartition(Partition{}, node)
			partitions = append(partitions, partition)
		}
	}
	return
}

func (f *partitionFinder) visitPartition(partition Partition, node Node) Partition {
	f.visited[node] = true
	partition = append(partition, node)
	for _, neighbor := range f.neighbors[node] {
		if !f.visited[neighbor] {
			partition = f.visitPartition(partition, neighbor)
		}
	}
	return partition
}
