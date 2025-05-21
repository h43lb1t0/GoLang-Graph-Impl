package graph

import (
	"fmt"
)

// DirectedGraph is a child of graph
type DirectedGraph struct {
	Graph
	edgeMap map[string]map[string]float64 // Map for O(1) edge lookups: from -> to -> length
}

func (graph *DirectedGraph) AddDirectedEdge(nodeId1, nodeId2 string, length float64) {
	if graph.edgeMap == nil {
		graph.edgeMap = make(map[string]map[string]float64)
	}

	if graph.containsVertex(graph.vertices, nodeId1) && graph.containsVertex(graph.vertices, nodeId2) {
		fromVertex := graph.getVertex(nodeId1)
		toVertex := graph.getVertex(nodeId2)

		// Initialize inner map if needed
		if _, exists := graph.edgeMap[nodeId1]; !exists {
			graph.edgeMap[nodeId1] = make(map[string]float64)
		}

		// Check if edge exists using map
		if _, exists := graph.edgeMap[nodeId1][nodeId2]; !exists {
			fromVertex.adjacent = append(fromVertex.adjacent, Edge{vertex: toVertex, length: length})
			graph.edgeMap[nodeId1][nodeId2] = length
		} else {
			fmt.Printf("Edge from %v to %v is already added\n", nodeId1, nodeId2)
		}
	} else {
		fmt.Printf("One of the vertices does not exist\n")
	}
}
