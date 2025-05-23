package main

import "AbgabeAlgo/graph"

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

// Print the shortest path from the webgraph
func webgraph() {
	graph := graph.InitWebgraph()
	graph.PrintDijkstraDistances("1")
}

// Main function with different tasks to uncomment and run
func main() {

	//problem98Challange()
	webgraph()

}
