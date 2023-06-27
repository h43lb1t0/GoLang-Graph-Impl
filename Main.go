package main

import "AbgabeAlgo/graph"

func main() {

	//test := &graph.Graph{}
	//directedGraph := &graph.DirectedGraph{}
	//unDirectedGraph := &graph.UnDirectedGraph{}

	/* test.AddVertex("A")
	test.AddVertex("B")
	test.AddVertex("C")
	test.AddVertex("D")
	test.AddVertex("E")
	test.AddVertex("E") */

	/* directedGraph.AddVertex("A")
	directedGraph.AddVertex("B")
	directedGraph.AddVertex("C")
	directedGraph.AddVertex("D")
	directedGraph.AddVertex("E") */

	/* unDirectedGraph.AddVertex("A")
	unDirectedGraph.AddVertex("B")
	unDirectedGraph.AddVertex("C")
	unDirectedGraph.AddVertex("D")
	unDirectedGraph.AddVertex("E")
	unDirectedGraph.AddVertex("F")

	unDirectedGraph.AddUndirectedEdge("A", "B", 1)
	unDirectedGraph.AddUndirectedEdge("A", "C", 5)
	unDirectedGraph.AddUndirectedEdge("D", "A", 2)
	unDirectedGraph.AddUndirectedEdge("B", "C", 1)
	unDirectedGraph.AddUndirectedEdge("B", "D", 2)
	unDirectedGraph.AddUndirectedEdge("E", "F", 4) */

	//unDirectedGraph.PrintDistances(unDirectedGraph.BFS("A"))

	/* unDirectedGraph.PrintAdjacent()

	unDirectedGraph.PrintUCC(unDirectedGraph.UCC()) */

	//test.PrintAdjacent()
	//test2.Graph.PrintAdjacent()

	//directedGraph.AddDirectedEdge("A", "B", 1)
	/* directedGraph.AddDirectedEdge("A", "B", 1)
	directedGraph.AddDirectedEdge("A", "C", 5)
	directedGraph.AddDirectedEdge("D", "A", 2)
	directedGraph.AddDirectedEdge("B", "C", 1)
	directedGraph.AddDirectedEdge("B", "D", 2) */

	//directedGraph.AddDirectedEdge("A", "H", 1)

	/* unDirectedGraph.AddUndirectedEdge("B", "C", 3)
	unDirectedGraph.AddUndirectedEdge("B", "D", 4) */

	/* fmt.Println("Directed Graph:")
	directedGraph.Graph.PrintAdjacent() */

	//directedGraph.Graph.PrintDistances(directedGraph.BFS("A"))
	/* bfs := directedGraph.BFS("A")

	directedGraph.PrintDistances(bfs) */

	/* dfs := directedGraph.DFS("A")

	directedGraph.PrintVisited(dfs) */

	/* fmt.Println("UnDirected Graph:")
	unDirectedGraph.Graph.PrintAdjacent()

	fmt.Printf("Number of Vertices in Directed Graph: %v\n", directedGraph.NumVertices())
	fmt.Printf("Number of Vertices in Undiredtec Graph: %v\n", unDirectedGraph.NumVertices())

	fmt.Printf("Number of Edges in Directed Graph: %v\n", directedGraph.NumEdges())
	fmt.Printf("Number of Edges in Undirected Graph: %v\n", unDirectedGraph.NumEdges()) */

	/* dag := &graph.DAG{}

	dag.AddVertex("A")
	dag.AddVertex("B")
	dag.AddVertex("C")
	dag.AddVertex("D")
	dag.AddVertex("E")

	dag.AddDirectedEdge("A", "B", 1)
	dag.AddDirectedEdge("B", "C", 1)
	dag.AddDirectedEdge("B", "D", 1)
	dag.AddDirectedEdge("B", "E", 1)
	dag.PrintAdjacent()

	dag.PrintTopoSort(dag.TopoSort()) */

	//directedGraph := graph.ParseGraphFromFile("./testdata/problem9.8.txt")

	//directedGraph.PrintAdjacent()

	//directedGraph.PrintDistances(directedGraph.BFS("35"))

	//directedGraph.PrintVisited(directedGraph.DFS("35"))

	/* directedGraph := graph.InitWebgraph()

	directedGraph.PrintDijkstraDistances("1") */

	/* graph := graph.DirectedGraph{}
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddDirectedEdge("2", "1", 1.0)
	graph.AddDirectedEdge("1", "2", 1.0)

	graph.PrintAdjacent()
	graph.PrintDijkstraDistances("1") */

	//directedGraph.PrintShortestPath("1", "7")

	//directedGraph.PrintAdjacent()

	dag := graph.DAG{}
	dag.AddVertex("1")
	dag.AddVertex("2")
	dag.AddVertex("3")
	dag.AddDirectedEdge("1", "2", 1.0)
	dag.AddDirectedEdge("2", "3", 1.0)

	dag.PrintTopoSort()

}
