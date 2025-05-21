// Package graph implements an adjacency list representation of a graph
// with support for both directed and undirected graphs.
package graph

import (
	"fmt"
	"math"
)

// Graph represents a graph using an adjacency list representation.
// It maintains a list of vertices and a map for O(1) vertex lookups.
type Graph struct {
	vertices  []*Vertex
	vertexMap map[string]*Vertex // Map for O(1) vertex lookups
}

// Edge represents a connection between two vertices with an associated length/weight.
type Edge struct {
	vertex *Vertex
	length float64
}

// Vertex represents a node in the graph with a unique key and its adjacent edges.
type Vertex struct {
	key      string
	adjacent []Edge
}

// getVertecies returns the list of all vertices in the graph.
func (graph *Graph) getVertecies() []*Vertex {
	return graph.vertices
}

// AddVertex adds a new vertex to the graph if it doesn't already exist.
// The vertex is identified by a unique string key.
func (graph *Graph) AddVertex(key string) {
	if graph.vertexMap == nil {
		graph.vertexMap = make(map[string]*Vertex)
	}

	if _, exists := graph.vertexMap[key]; !exists {
		vertex := &Vertex{key: key}
		graph.vertices = append(graph.vertices, vertex)
		graph.vertexMap[key] = vertex
	}
}

// containsVertex checks if a vertex with the given key exists in the graph.
// Returns true if the vertex exists, false otherwise.
func (graph *Graph) containsVertex(s []*Vertex, k string) bool {
	if graph.vertexMap == nil {
		return false
	}
	_, exists := graph.vertexMap[k]
	return exists
}

// containsEdge checks if an edge exists between the given vertex and its adjacent vertices.
// Returns true if the edge exists, false otherwise.
func (graph *Graph) containsEdge(edges []Edge, k string) bool {
	for _, edge := range edges {
		if k == edge.vertex.key {
			return true
		}
	}
	return false
}

// getVertex returns the vertex with the given key if it exists.
// Returns nil if no such vertex exists.
func (graph *Graph) getVertex(k string) *Vertex {
	if graph.vertexMap == nil {
		return nil
	}
	return graph.vertexMap[k]
}

// getEdge returns the edge to the vertex with the given key if it exists.
// Returns nil if no such edge exists.
func (graph *Graph) getEdge(k string) *Edge {
	for _, v := range graph.vertices {
		for _, edge := range v.adjacent {
			if k == edge.vertex.key {
				return &edge
			}
		}
	}
	return nil
}

// NumVertices returns the total number of vertices in the graph.
func (graph *Graph) NumVertices() int {
	return len(graph.vertices)
}

// NumEdges returns the total number of edges in the graph.
func (graph *Graph) NumEdges() int {
	count := 0
	for _, v := range graph.vertices {
		count += len(v.adjacent)
	}
	return count
}

/* func (graph *Graph) IsDirected() bool {

} */

// BFS performs a breadth-first search starting from the given node.
// Returns a map of node IDs to their distances from the start node.
// The distance represents the number of edges in the shortest path.
func (graph *Graph) BFS(nodeId string) map[string]int {
	visited := make(map[string]bool)
	for _, v := range graph.vertices {
		visited[v.key] = false
	}
	return graph.BFS_intern(nodeId, visited)
}

// BFS_intern is the internal implementation of BFS that takes a visited map.
// Returns a map of node IDs to their distances from the start node.
func (graph *Graph) BFS_intern(nodeId string, visited map[string]bool) map[string]int {

	//fmt.Println("BFS:")

	if !graph.containsVertex(graph.vertices, nodeId) {
		fmt.Printf("Node %v does not exist\n", nodeId)
		return nil
	}

	distances := make(map[string]int)

	visited[nodeId] = true

	// a new queue
	queue := make([]*Vertex, 0)

	// while the queue is not empty
	queue = append(queue, graph.getVertex(nodeId))
	distances[nodeId] = 0
	for len(queue) > 0 {

		// remove the first element from the queue
		v := queue[0]
		queue = queue[1:]

		for _, w := range v.adjacent {
			if !visited[w.vertex.key] {
				queue = append(queue, w.vertex)
				visited[w.vertex.key] = true
				distances[w.vertex.key] = distances[v.key] + 1
			}
		}
	}

	return distances

}

// DFS performs a depth-first search starting from the given node.
// Returns a map of node IDs to boolean values indicating if they were visited.
func (graph *Graph) DFS(nodeId string) map[string]bool {
	//fmt.Println("DFS:")

	visited := make(map[string]bool)

	if !graph.containsVertex(graph.vertices, nodeId) {
		fmt.Printf("Node %v does not exist\n", nodeId)
		return nil
	}

	for _, v := range graph.vertices {
		visited[v.key] = false
	}
	visited[nodeId] = true

	v := graph.getVertex(nodeId)

	stack := &Stack{}

	stack.Push(v)

	for !stack.IsEmpty() {
		v = stack.Pop()
		for _, w := range v.adjacent {
			if !visited[w.vertex.key] {
				stack.Push(w.vertex)
				visited[w.vertex.key] = true
			}
		}
	}

	return visited
}

// PrintAdjacent prints the adjacency list representation of the graph.
func (graph *Graph) PrintAdjacent() {
	for _, v := range graph.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, e := range v.adjacent {
			fmt.Printf("(%v, %v) ", e.vertex.key, e.length)
		}
	}
	fmt.Println()
}

// PrintDistances prints the distances from a start node to all other nodes.
func (graph *Graph) PrintDistances(distances map[string]int) {
	fmt.Println("Distances:")
	for k, v := range distances {
		fmt.Printf("Node %v: %v\n", k, v)
	}
}

// PrintVisited prints the visited status of all nodes after a traversal.
func (graph *Graph) PrintVisited(visited map[string]bool) {
	fmt.Println("Visited:")
	for k, v := range visited {
		fmt.Printf("Node %v: %v\n", k, v)
	}
}

// Dijkstra performs Dijkstra's shortest path algorithm starting from the given node.
// Returns a map of node IDs to their distances from the start node.
// The distance represents the length of the shortest path.
func (graph *Graph) Dijkstra(start string) map[string]float64 {
	if !graph.containsVertex(graph.vertices, start) {
		fmt.Printf("Start node %v does not exist\n", start)
		return nil
	}

	distances := make(map[string]float64)
	visited := make(map[string]bool)

	// Initialize distances to infinity and visited to false
	for _, vertex := range graph.vertices {
		distances[vertex.key] = math.Inf(1)
		visited[vertex.key] = false
	}

	// Set start node distance to 0
	distances[start] = 0

	// Create and initialize min heap
	heap := newMinHeap(graph.NumVertices())
	heap.enqueue(&HeapNode{NodeId: start, Distance: 0})

	// Main Dijkstra loop
	for len(heap.nodes) > 0 {
		minHeapNode := heap.dequeue()
		currentNode := minHeapNode.NodeId

		if !visited[currentNode] {
			visited[currentNode] = true

			// Process all adjacent vertices
			for _, edge := range graph.getVertex(currentNode).adjacent {
				neighbor := edge.vertex.key
				if !visited[neighbor] {
					newDistance := distances[currentNode] + edge.length
					if newDistance < distances[neighbor] {
						distances[neighbor] = newDistance
						heap.enqueue(&HeapNode{NodeId: neighbor, Distance: newDistance})
					}
				}
			}
		}
	}

	return distances
}

// PrintDijkstraDistances prints the distances from a start node to all other nodes.
func (graph *Graph) PrintDijkstraDistances(start string) {
	distances := graph.Dijkstra(start)
	if distances != nil {
		for node, distance := range distances {
			if math.IsInf(distance, 1) {
				fmt.Printf("From %v to %v: INFINITY\n", start, node)
			} else {
				fmt.Printf("From %v to %v: Distance %v\n", start, node, distance)
			}
		}
	} else {
		fmt.Println("Start node does not exist, no distances to print.")
	}
}
