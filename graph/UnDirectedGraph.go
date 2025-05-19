package graph

import "fmt"

type UnDirectedGraph struct {
	Graph
	edgeMap map[string]map[string]float64 // Map for O(1) edge lookups: from -> to -> length
}

func (graph *UnDirectedGraph) AddUndirectedEdge(nodeId1, nodeId2 string, length float64) {
	if graph.edgeMap == nil {
		graph.edgeMap = make(map[string]map[string]float64)
	}

	if graph.containsVertex(graph.vertices, nodeId1) && graph.containsVertex(graph.vertices, nodeId2) {
		fromVertex := graph.getVertex(nodeId1)
		toVertex := graph.getVertex(nodeId2)

		// Initialize inner maps if needed
		if _, exists := graph.edgeMap[nodeId1]; !exists {
			graph.edgeMap[nodeId1] = make(map[string]float64)
		}
		if _, exists := graph.edgeMap[nodeId2]; !exists {
			graph.edgeMap[nodeId2] = make(map[string]float64)
		}

		// Check if edge exists using map
		if _, exists := graph.edgeMap[nodeId1][nodeId2]; !exists {
			fromVertex.adjacent = append(fromVertex.adjacent, Edge{vertex: toVertex, length: length})
			toVertex.adjacent = append(toVertex.adjacent, Edge{vertex: fromVertex, length: length})
			graph.edgeMap[nodeId1][nodeId2] = length
			graph.edgeMap[nodeId2][nodeId1] = length
		} else {
			fmt.Printf("Edge between %v and %v is already added\n", nodeId1, nodeId2)
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
