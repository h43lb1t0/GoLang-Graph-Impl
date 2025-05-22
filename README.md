# GoLang-Graph-Impl

## Description
This project implements various graph data structures and algorithms in Go as part of a university assignment. It provides implementations for both directed and undirected graphs, along with common graph algorithms.

## Features
- **Graph Types**:
  - Directed Graph
  - Undirected Graph
  - Directed Acyclic Graph (DAG)

- **Core Graph Operations**:
  - Add vertices and edges
  - Count vertices and edges
  - Check vertex/edge existence
  - Print adjacency lists

- **Graph Algorithms**:
  - Breadth-First Search (BFS)
  - Depth-First Search (DFS)
  - Dijkstra's Shortest Path Algorithm
  - Topological Sort (for DAGs)
  - Undirected Connected Components (UCC)

## Project Structure
```
.
├── graph/                  # Core graph implementation
│   ├── Graph.go           # Base graph structure and common operations
│   ├── DirectedGraph.go   # Directed graph implementation
│   ├── UnDirectedGraph.go # Undirected graph implementation
│   ├── dag.go             # Directed Acyclic Graph implementation
│   ├── heap.go            # Min heap implementation for Dijkstra's algorithm
│   ├── stack.go           # Stack implementation for DFS
│   └── interfaces.go      # Graph interfaces
├── tests/                 # Test files
│   ├── DirectedGraph_test.go
│   ├── UndirectedGraph_test.go
│   └── dag_test.go
└── Main.go               # Example usage and problem solutions
```

## Usage
The project includes example implementations for:
1. Problem 9.8 - Shortest path calculations
2. Web graph analysis using Google's web graph dataset

To run the examples:
```bash
go run Main.go
```

## Testing
Run the test suite using:
```bash
go test ./tests/...
```

## Implementation Details
- Uses adjacency list representation for efficient graph operations
- Implements O(1) vertex and edge lookups using maps
- Provides both directed and undirected graph variants
- Includes comprehensive test coverage for all graph operations and algorithms
