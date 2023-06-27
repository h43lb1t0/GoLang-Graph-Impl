// Package graph implements an adjacency list for a graph representation
// and some selected algorithms on such graphs.
package graph

// Interface of a Graph
// - each node should have an unique nodeId (externally set) of type string
/* type Graph interface {

	// add a vertex. If the vertex already exists in the graph: do nothing
	AddVertex(nodeId string) //, data interface{})

	// get some generic meta data of the graph:
	NumVertices() int // number of vertices
	NumEdges() int    // number of edges
	//IsDirected() bool // is a directed or undirected graph?

	// some basic algorithms:
	// the returned map maps for each reachable
	// node the distance in layers from given nodeId to each node
	BFS(nodeId string) map[string]int

	// maps the reachable node to true
	DFS(nodeId string) map[string]bool
} */

/* type DirectedGraph interface {
	Graph

	//
	Dijkstra(id string) map[string]float64

	//HasVertex(nodeId string) bool
	// the vertices must exists
	AddDirectedEdge(nodeId1, nodeId2 string, length float64) //, data interface{})

	//BellmannFord(id string) ([]float64, bool)
	//FloydWarshall() ([][]float64, bool) // no negative cycles
} */

// directed acyclic graph
/* type DAG interface {
	DirectedGraph
	// algorithms
	// TopoSort returns an map from nodeId to topological order
	TopoSort() map[string]int
} */

/* type UnDirectedGraph interface {
	Graph
	AddUndirectedEdge(nodeId1, nodeId2 string, length float64)

	// undirected connected components
	// return a map from nodeID to connected component ID
	UCC() map[string]int

	//Kruskal(graph *AdjacencyList) (*AdjacencyList, float64)
} */
