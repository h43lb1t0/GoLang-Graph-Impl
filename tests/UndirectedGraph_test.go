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
