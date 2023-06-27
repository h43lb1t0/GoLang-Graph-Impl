package graph

import "fmt"

type DAG struct {
	DirectedGraph
}

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

func (graph *DAG) PrintTopoSort(ordering map[string]int) {
	fmt.Println("Topological Sort:")
	for k, v := range ordering {
		fmt.Printf("%v: %v\n", k, v)
	}
}
