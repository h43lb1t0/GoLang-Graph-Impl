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
	graph.AddDirectedEdge("1", "2", 1.0)

	if graph.NumEdges() != 1 {
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
	graph.AddDirectedEdge("1", "2", 1.0)
	graph.AddVertex("3")
	Distances := graph.Dijkstra("1")

	expectedDistances := make(map[string]float64)
	expectedDistances["1"] = 0.0
	expectedDistances["2"] = 1.0
	expectedDistances["3"] = math.Inf(1)

	if len(Distances) != len(expectedDistances) && len(Distances) > 0 {
		for id, d := range Distances {
			if expectedDistances[id] != d {
				t.Errorf("Dijkstra method error")
			}
		}
	} else {
		t.Errorf("Dijkstra method error")

	}
}
