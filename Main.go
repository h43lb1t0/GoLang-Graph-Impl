package main

import (
	"AbgabeAlgo/graph"
	"fmt"
)

// Print the shortest path from the task
func problem98Challange() {
	var endNodes = []string{
		"7", "37", "59", "82", "99", "115", "133", "165", "188", "197",
	}
	graph := graph.InitProblem98()

	for _, endNode := range endNodes {
		graph.PrintDijkstraFromTo("1", endNode)
	}
}

func problem98test() {
	graph := graph.InitProblem98test()
	graph.PrintDijkstraDistances("1")
}

func problem98testDFS() {
	graph := graph.InitProblem98test()

	result := graph.DFS("1")

	for k, v := range result {
		fmt.Printf("Node %s: %v\n", k, v)
	}
}

// Print the shortest path from the webgraph
func webgraph() {
	graph := graph.InitWebgraph()
	graph.PrintDijkstraDistances("1")
}

// Main function with different tasks to uncomment and run
func main() {

	//problem98Challange()
	//problem98test()
	//problem98testDFS()
	webgraph()
}
