package tests

import (
	"AbgabeAlgo/graph"
	"math"
	"testing"
)

func TestAddVertexToDirectedGraph(t *testing.T) {
	graph := graph.DirectedGraph{}
	graph.AddVertex("1")
	if graph.NumVertices() != 1 {
		t.Errorf("Failed to add a vertex")
	}
}

func TestAddDirectedEdge(t *testing.T) {
	graph := graph.DirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddDirectedEdge("1", "2", 1.0)

	if graph.NumEdges() != 1 {
		t.Errorf("Failed to add an undirected edge")
	}
}

func TestNumVerticesOnDirectedGraph(t *testing.T) {
	graph := graph.DirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	if graph.NumVertices() != 2 {
		t.Errorf("NumVertices method error")
	}
}

func TestNumEdgesOnDirectedGraph(t *testing.T) {
	graph := graph.DirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddDirectedEdge("1", "2", 1.0)
	graph.AddDirectedEdge("1", "3", 2.0)

	if graph.NumEdges() != 2 {
		t.Errorf("NumEdges method error")
	}
}

func TestBFSOnDirectedGraph(t *testing.T) {
	graph := graph.DirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddDirectedEdge("1", "2", 1.0)
	distances := graph.BFS("1")

	// create a new map with the expected distances
	expectedDistances := make(map[string]int)
	expectedDistances["1"] = 0.0
	expectedDistances["2"] = 1.0

	// check if distances and expectedDistances are equal
	if len(distances) == len(expectedDistances) {
		for k, v := range distances {
			if expectedDistances[k] != v {
				t.Errorf("BFS method error")
			}
		}
	} else {
		t.Errorf("BFS method error")
	}

}

func TestDFSOnDirectedGraph(t *testing.T) {
	graph := graph.DirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddDirectedEdge("1", "2", 1.0)
	visited := graph.DFS("1")

	// create a new map with the expected visited nodes
	expectedVisited := make(map[string]bool)
	expectedVisited["1"] = true
	expectedVisited["2"] = true

	// check if visited and expectedVisited are equal
	if len(visited) == len(expectedVisited) {
		for k, v := range visited {
			if expectedVisited[k] != v {
				t.Errorf("DFS method error")
			}
		}
	} else {
		t.Errorf("DFS method error")
	}
}

func TestDijkstra(t *testing.T) {
	graph := graph.DirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddDirectedEdge("2", "1", 1.0)
	graph.AddDirectedEdge("1", "2", 2.0)
	graph.AddVertex("3")
	Distances := graph.Dijkstra("1")

	expectedDistances := make(map[string]float64)
	expectedDistances["1"] = 0.0
	expectedDistances["2"] = 2.0
	expectedDistances["3"] = math.Inf(1)

	if len(Distances) == len(expectedDistances) {
		for id, d := range Distances {
			if expectedDistances[id] != d {
				t.Errorf("Dijkstra method error: expected distance %v for node %v, got %v", expectedDistances[id], id, d)
			}
		}
	} else {
		t.Errorf("Dijkstra method error: expected %d distances, got %d", len(expectedDistances), len(Distances))
	}
}

func TestDijkstraComplex(t *testing.T) {
	graph := graph.DirectedGraph{}

	// Create a more complex directed graph
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddVertex("F")

	// Add edges with different weights
	graph.AddDirectedEdge("A", "B", 4.0)
	graph.AddDirectedEdge("A", "C", 2.0)
	graph.AddDirectedEdge("B", "C", 1.0)
	graph.AddDirectedEdge("B", "D", 5.0)
	graph.AddDirectedEdge("C", "D", 8.0)
	graph.AddDirectedEdge("C", "E", 10.0)
	graph.AddDirectedEdge("D", "E", 2.0)
	graph.AddDirectedEdge("D", "F", 6.0)
	graph.AddDirectedEdge("E", "F", 3.0)

	// Calculate shortest paths from vertex A
	distances := graph.Dijkstra("A")

	// Expected shortest path distances from A
	expectedDistances := make(map[string]float64)
	expectedDistances["A"] = 0.0  // Distance to self
	expectedDistances["B"] = 4.0  // A -> B
	expectedDistances["C"] = 2.0  // A -> C
	expectedDistances["D"] = 9.0  // A -> B -> D
	expectedDistances["E"] = 11.0 // A -> B -> D -> E
	expectedDistances["F"] = 14.0 // A -> B -> D -> E -> F

	if len(distances) == len(expectedDistances) {
		for id, d := range distances {
			if expectedDistances[id] != d {
				t.Errorf("Dijkstra method error: expected distance %v for node %v, got %v", expectedDistances[id], id, d)
			}
		}
	} else {
		t.Errorf("Dijkstra method error: expected %d distances, got %d", len(expectedDistances), len(distances))
	}
}
