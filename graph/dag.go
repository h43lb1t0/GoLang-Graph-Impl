package graph

import "fmt"

// DAG represents a Directed Acyclic Graph, which is a directed graph with no directed cycles.
// It extends the DirectedGraph type and provides topological sorting functionality.
type DAG struct {
	DirectedGraph
}

// TopoSort performs a topological sort on the DAG and returns a map where:
// - Keys are vertex identifiers
// - Values are their topological ordering (1 to n, where n is the number of vertices)
// The ordering represents a linear ordering of vertices where all edges point forward.
func (graph *DAG) TopoSort() map[string]int {
	curLabel := graph.NumVertices()
	ordering := make(map[string]int)
	visited := make(map[string]bool)
	for _, v := range graph.vertices {
		visited[v.key] = false
	}
	for _, v := range graph.vertices {
		if !visited[v.key] {
			graph.DfsTopoSort(v, visited, &curLabel, ordering)
		}
	}

	return ordering
}

// DfsTopoSort is a helper function that performs a depth-first search to build the topological ordering.
// It recursively visits all unvisited neighbors of a vertex and assigns topological labels in reverse order.
// Parameters:
//   - vertex: The current vertex being processed
//   - visited: Map tracking visited vertices
//   - curLabel: Pointer to the current label being assigned
//   - ordering: Map storing the final topological ordering
func (graph *DAG) DfsTopoSort(vertex *Vertex, visited map[string]bool, curLabel *int, ordering map[string]int) {
	visited[vertex.key] = true
	for _, edge := range vertex.adjacent {
		if !visited[edge.vertex.key] {
			graph.DfsTopoSort(edge.vertex, visited, curLabel, ordering)
		}
	}
	ordering[vertex.key] = *curLabel
	*curLabel = *curLabel - 1
}

// PrintTopoSort prints the topological ordering of the DAG to stdout.
// The output shows each vertex and its corresponding topological order.
func (graph *DAG) PrintTopoSort() {
	ordering := graph.TopoSort()
	fmt.Println("Topological Sort:")
	for k, v := range ordering {
		fmt.Printf("%v: %v\n", k, v)
	}
}
