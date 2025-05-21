package tests

import (
	"AbgabeAlgo/graph"
	"testing"
)

func TestAddVertex(t *testing.T) {
	graph := graph.UnDirectedGraph{}
	graph.AddVertex("1")
	if graph.NumVertices() != 1 {
		t.Errorf("Failed to add a vertex")
	}
}

func TestAddUndirectedEdge(t *testing.T) {
	graph := graph.UnDirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddUndirectedEdge("1", "2", 1.0)

	if graph.NumEdges() != 2 {
		t.Errorf("Failed to add an undirected edge")
	}
}

func TestNumVertices(t *testing.T) {
	graph := graph.UnDirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	if graph.NumVertices() != 2 {
		t.Errorf("NumVertices method error")
	}
}

func TestNumEdges(t *testing.T) {
	graph := graph.UnDirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddUndirectedEdge("1", "2", 1.0)

	if graph.NumEdges() != 2 {
		t.Errorf("NumEdges method error")
	}
}

func TestBFS(t *testing.T) {
	graph := graph.UnDirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddUndirectedEdge("1", "2", 1.0)
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

func TestDFS(t *testing.T) {
	graph := graph.UnDirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddUndirectedEdge("1", "2", 1.0)
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

func TestUCC(t *testing.T) {
	graph := graph.UnDirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddUndirectedEdge("1", "2", 1.0)
	graph.AddVertex("3")
	components := graph.UCC()

	if len(components) != 3 || components["3"] == components["1"] {
		t.Errorf("UCC method error")
	}
}

func TestDijkstraUndirected(t *testing.T) {
	graph := graph.UnDirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddUndirectedEdge("1", "2", 2.0)
	graph.AddUndirectedEdge("2", "3", 3.0)
	graph.AddUndirectedEdge("1", "3", 6.0)

	distances := graph.Dijkstra("1")

	expectedDistances := make(map[string]float64)
	expectedDistances["1"] = 0.0
	expectedDistances["2"] = 2.0
	expectedDistances["3"] = 5.0 // Should go through node 2

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

func TestDijkstraUndirectedComplex(t *testing.T) {
	graph := graph.UnDirectedGraph{}

	// Create a more complex undirected graph
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddVertex("F")

	// Add edges with different weights
	graph.AddUndirectedEdge("A", "B", 4.0)
	graph.AddUndirectedEdge("A", "C", 2.0)
	graph.AddUndirectedEdge("B", "C", 1.0)
	graph.AddUndirectedEdge("B", "D", 5.0)
	graph.AddUndirectedEdge("C", "D", 8.0)
	graph.AddUndirectedEdge("C", "E", 10.0)
	graph.AddUndirectedEdge("D", "E", 2.0)
	graph.AddUndirectedEdge("D", "F", 6.0)
	graph.AddUndirectedEdge("E", "F", 3.0)

	// Calculate shortest paths from vertex A
	distances := graph.Dijkstra("A")

	// Expected shortest path distances from A
	expectedDistances := make(map[string]float64)
	expectedDistances["A"] = 0.0  // Distance to self
	expectedDistances["B"] = 3.0  // A -> C -> B
	expectedDistances["C"] = 2.0  // A -> C
	expectedDistances["D"] = 8.0  // A -> C -> B -> D
	expectedDistances["E"] = 10.0 // A -> C -> B -> D -> E
	expectedDistances["F"] = 13.0 // A -> C -> B -> D -> E -> F

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
