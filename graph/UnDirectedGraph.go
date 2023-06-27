package graph

import "fmt"

type UnDirectedGraph struct {
	Graph
}

func (graph *UnDirectedGraph) AddUndirectedEdge(nodeId1, nodeId2 string, length float64) {
	if graph.containsVertex(graph.vertices, nodeId1) && graph.containsVertex(graph.vertices, nodeId2) {
		fromVertex := graph.getVertex(nodeId1)
		toVertex := graph.getVertex(nodeId2)
		if !graph.containsEdge(fromVertex.adjacent, nodeId2) {
			fromVertex.adjacent = append(fromVertex.adjacent, Edge{vertex: toVertex, length: length})
			toVertex.adjacent = append(toVertex.adjacent, Edge{vertex: fromVertex, length: length})
		} else {
			fmt.Printf("Edge from %v to %v is already added\n", nodeId1, nodeId2)
		}
	} else {
		fmt.Printf("One of the vertices does not exist\n")
	}
}

func (graph *UnDirectedGraph) UCC() map[string]int {
	components := make(map[string]int)
	commpId := 0

	for _, v := range graph.vertices {
		if components[v.key] == 0 {
			visited := make(map[string]bool) // Create a new visited map for each BFS call
			graph.BFS_intern(v.key, visited)

			for k := range visited {
				if components[k] == 0 {
					components[k] = commpId
				}
			}
			commpId++
		}
	}

	return components
}

func (graph *UnDirectedGraph) PrintUCC(components map[string]int) {
	fmt.Println("Connected Components:")
	for k, v := range components {
		fmt.Printf("Vertex %v is in component %v\n", k, v)
	}
}
