package graph

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Edge struct {
	vertex *Vertex
	length float64
}

type Vertex struct {
	key      string
	adjacent []Edge
}

func (graph *Graph) getVertecies() []*Vertex {
	return graph.vertices
}

func (graph *Graph) AddVertex(key string) {
	if !graph.containsVertex(graph.vertices, key) {
		graph.vertices = append(graph.vertices, &Vertex{key: key})
	} else {
		fmt.Printf("Vertex %v is already added\n", key)
	}
}

func (graph *Graph) containsVertex(s []*Vertex, k string) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (graph *Graph) containsEdge(edges []Edge, k string) bool {
	for _, edge := range edges {
		if k == edge.vertex.key {
			return true
		}
	}
	return false
}

func (graph *Graph) getVertex(k string) *Vertex {
	for _, v := range graph.vertices {
		if k == v.key {
			return v
		}
	}
	return nil
}

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

func (graph *Graph) NumVertices() int {
	return len(graph.vertices)
}

func (graph *Graph) NumEdges() int {
	count := 0
	for _, v := range graph.vertices {
		count += len(v.adjacent)
	}
	return count
}

/* func (graph *Graph) IsDirected() bool {

} */

func (graph *Graph) BFS(nodeId string) map[string]int {
	visited := make(map[string]bool)
	for _, v := range graph.vertices {
		visited[v.key] = false
	}
	return graph.BFS_intern(nodeId, visited)
}

// the returned map maps for each reachable
// node the distance in layers from given nodeId to each node
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

func (graph *Graph) PrintAdjacent() {
	for _, v := range graph.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, e := range v.adjacent {
			fmt.Printf("(%v, %v) ", e.vertex.key, e.length)
		}
	}
	fmt.Println()
}

func (graph *Graph) PrintDistances(distances map[string]int) {
	fmt.Println("Distances:")
	for k, v := range distances {
		fmt.Printf("Node %v: %v\n", k, v)
	}
}

func (graph *Graph) PrintVisited(visited map[string]bool) {
	fmt.Println("Visited:")
	for k, v := range visited {
		fmt.Printf("Node %v: %v\n", k, v)
	}
}
