package graph

import (
	"fmt"
)

// DirectedGraph represents a directed graph data structure where edges have a specific direction
// from one vertex to another. It extends the base Graph type and maintains an additional
// edgeMap for O(1) edge lookups.
type DirectedGraph struct {
	Graph
	edgeMap map[string]map[string]float64 // Map for O(1) edge lookups: from -> to -> length
}

// AddDirectedEdge adds a directed edge from nodeId1 to nodeId2 with the specified length.
// The edge will only be added if both vertices exist in the graph and if the edge doesn't
// already exist. If the edge already exists, a message will be printed.
//
// Parameters:
//   - nodeId1: The ID of the source vertex
//   - nodeId2: The ID of the target vertex
//   - length: The length/weight of the edge
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
