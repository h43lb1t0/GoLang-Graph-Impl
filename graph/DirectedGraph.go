package graph

import (
	"fmt"
	"math"
)

// DirectedGraph is a child of graph
type DirectedGraph struct {
	Graph
}

func (graph *DirectedGraph) AddDirectedEdge(nodeId1, nodeId2 string, length float64) {
	if graph.containsVertex(graph.vertices, nodeId1) && graph.containsVertex(graph.vertices, nodeId2) {
		fromVertex := graph.getVertex(nodeId1)
		toVertex := graph.getVertex(nodeId2)
		if !graph.containsEdge(fromVertex.adjacent, nodeId2) {
			fromVertex.adjacent = append(fromVertex.adjacent, Edge{vertex: toVertex, length: length})
		} else {
			fmt.Printf("Edge from %v to %v is already added\n", nodeId1, nodeId2)
		}
	} else {
		fmt.Printf("One of the vertices does not exist\n")
	}
}

func (graph *DirectedGraph) Dijkstra(start string) map[string]float64 {

	if graph.containsVertex(graph.vertices, start) {

		Distances := make(map[string]float64)

		visted := make(map[string]bool)

		for _, vertex := range graph.vertices {
			Distances[vertex.key] = math.Inf(1)
			visted[vertex.key] = false
		}

		Distances[start] = 0

		heap := newMinHeap(graph.NumVertices())
		heap.enqueue(&HeapNode{NodeId: start, Distance: 0})

		for len(heap.nodes) > 0 {
			minHeapNode := heap.dequeue()
			currentHeapNode := minHeapNode.NodeId
			if !visted[currentHeapNode] {
				visted[currentHeapNode] = true
				for _, edge := range graph.getVertex(currentHeapNode).adjacent {
					neighbour := edge.vertex.key
					if !visted[neighbour] {
						newDistance := Distances[currentHeapNode] + edge.length
						if newDistance < Distances[neighbour] {
							Distances[neighbour] = newDistance
							heap.enqueue(&HeapNode{NodeId: neighbour, Distance: newDistance})
						}

					}
				}
			}

		}

		return Distances
	}
	fmt.Printf("Start node %v does not exist\n", start)
	return nil
}

func (graph *DirectedGraph) PrintDijkstraDistances(start string) {
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
