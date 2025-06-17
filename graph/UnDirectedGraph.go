package graph

import "fmt"

// UnDirectedGraph represents an undirected graph data structure.
// It extends the base Graph type
type UnDirectedGraph struct {
	Graph
}

// AddUndirectedEdge adds an undirected edge between two vertices with the specified length.
// If either vertex doesn't exist or the edge already exists, appropriate messages are printed.
// The edge is added in both directions since this is an undirected graph.
//
// Parameters:
//   - nodeId1: The ID of the first vertex
//   - nodeId2: The ID of the second vertex
//   - length: The length/weight of the edge
func (graph *UnDirectedGraph) AddUndirectedEdge(nodeId1, nodeId2 string, length float64) {
	if graph.edgeMap == nil {
		graph.edgeMap = make(map[string]map[string]float64)
	}

	if graph.containsVertex(nodeId1) && graph.containsVertex(nodeId2) {
		fromVertex := graph.getVertex(nodeId1)
		toVertex := graph.getVertex(nodeId2)

		// Initialize inner maps if needed
		if _, exists := graph.edgeMap[nodeId1]; !exists {
			graph.edgeMap[nodeId1] = make(map[string]float64)
		}
		if _, exists := graph.edgeMap[nodeId2]; !exists {
			graph.edgeMap[nodeId2] = make(map[string]float64)
		}

		// Check if edge exists using map
		if _, exists := graph.edgeMap[nodeId1][nodeId2]; !exists {
			fromVertex.adjacent = append(fromVertex.adjacent, Edge{vertex: toVertex, length: length})
			toVertex.adjacent = append(toVertex.adjacent, Edge{vertex: fromVertex, length: length})
			graph.edgeMap[nodeId1][nodeId2] = length
			graph.edgeMap[nodeId2][nodeId1] = length
		} else {
			fmt.Printf("Edge between %v and %v is already added\n", nodeId1, nodeId2)
		}
	} else {
		fmt.Printf("One of the vertices does not exist\n")
	}
}

// UCC (Undirected Connected Components) identifies all connected components in the graph.
// It uses BFS to explore each component and assigns a unique component ID to each vertex.
//
// Returns:
//   - A map where keys are vertex IDs and values are their component IDs
func (graph *UnDirectedGraph) UCC() map[string]int {
	components := make(map[string]int)
	commpId := 0

	for _, v := range graph.vertices {
		if components[v.key] == 0 {
			visited := make(map[string]bool) // Create a new visited map for each BFS call
			graph.bfs_intern(v.key, visited)

			for k := range visited {
				if components[k] == 0 {
					components[k] = commpId
				}
			}
			commpId++
		}
	}

	return components
}

// PrintUCC prints the connected components of the graph in a human-readable format.
// It displays which vertices belong to which components.
//
// Parameters:
//   - components: A map containing the component assignments for each vertex
func (graph *UnDirectedGraph) PrintUCC(components map[string]int) {
	fmt.Println("Connected Components:")
	for k, v := range components {
		fmt.Printf("Vertex %v is in component %v\n", k, v)
	}
}
