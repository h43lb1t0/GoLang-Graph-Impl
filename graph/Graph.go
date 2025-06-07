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
func (graph *Graph) Dijkstra(startNodeId string) map[string]float64 {
	if graph.getVertex(startNodeId) == nil {
		fmt.Printf("start node %q does not exist in the graph", startNodeId)
	}

	distances := make(map[string]float64)
	visited := make(map[string]bool) // Tracks nodes whose shortest path is finalized

	// Initialize distances: Inf for all, 0 for startNodeId.
	for _, vertex := range graph.vertices {
		distances[vertex.key] = math.Inf(1)
	}
	distances[startNodeId] = 0

	// Create and initialize min heap.
	pq := NewMinHeap(graph.NumVertices())
	pq.Enqueue(&HeapNode{NodeId: startNodeId, Distance: 0})

	// Main Dijkstra loop
	for !pq.IsEmpty() {
		// Extract node with the smallest distance from the priority queue.
		minNode := pq.Dequeue()
		u_id := minNode.NodeId
		// u_distance := minNode.Distance // This is distances[u_id]

		// If u has already been finalized, skip (shouldn't happen often with decreaseKey).
		if visited[u_id] {
			fmt.Printf("Node %s is already added", u_id)
			continue
		}
		visited[u_id] = true // Mark u as finalized.

		u_vertex := graph.getVertex(u_id)
		// This check is more for internal consistency; u_id came from the graph.
		if u_vertex == nil {
			fmt.Printf("internal error: dequeued node %s not found via getVertex", u_id)
		}

		// Process all adjacent vertices (neighbors) of u.
		for _, edge := range u_vertex.adjacent {
			v_id := edge.vertex.key
			weight_uv := edge.length

			// If v is already finalized, we can't find a shorter path to it through u.
			if visited[v_id] {
				continue
			}

			// Relax edge (u,v):
			// If a shorter path to v is found through u.
			if distances[u_id]+weight_uv < distances[v_id] {
				newDistanceToV := distances[u_id] + weight_uv
				distances[v_id] = newDistanceToV

				if pq.Contains(v_id) {
					// v is already in the priority queue, update its distance.
					err := pq.DecreaseKey(v_id, newDistanceToV)
					if err != nil {
						// This might indicate an issue with heap logic or unexpected state.
						// For example, if newDistanceToV was not actually smaller than
						// the value stored *within the heap structure itself* for v_id.
						// This shouldn't happen if distances map and heap are consistent.
						fmt.Printf("Dijkstra: Info/Warning from decreaseKey for %s: %v\n", v_id, err)
					}
				} else {
					// v is not in the priority queue, add it.
					pq.Enqueue(&HeapNode{NodeId: v_id, Distance: newDistanceToV})
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

// PrintDijkstraFromTo prints the shortest path distance between two specific nodes.
// If no path exists, it prints "NO PATH". If the nodes don't exist, it prints nothing.
func (graph *Graph) PrintDijkstraFromTo(start string, end string) {
	distances := graph.Dijkstra(start)
	if distances != nil {
		if math.IsInf(distances[end], 1) {
			fmt.Printf("From %v to %v: NO PATH\n", start, end)
		} else {
			fmt.Printf("From %v to %v: Distance %v\n", start, end, distances[end])
		}
	}
}
