package main

import "AbgabeAlgo/graph"

func main() {

	directedGraph := graph.InitWebgraph()

	directedGraph.PrintDijkstraDistances("1")

}
