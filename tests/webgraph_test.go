package tests

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"

	"AbgabeAlgo/graph"
)

// loadGonumGraph loads the webgraph into a gonum graph structure
func loadGonumGraph(filename string) *simple.DirectedGraph {
	g := simple.NewDirectedGraph()

	// Get the absolute path to the testdata directory
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// If we're in the tests directory, go up one level
	if filepath.Base(wd) == "tests" {
		wd = filepath.Dir(wd)
	}
	filepath := filepath.Join(wd, "testdata", filepath.Base(filename))

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}

		fromID := int64(parseInt(fields[0]))
		toID := int64(parseInt(fields[1]))

		if g.Node(fromID) == nil {
			g.AddNode(simple.Node(fromID))
		}
		if g.Node(toID) == nil {
			g.AddNode(simple.Node(toID))
		}
		g.SetEdge(g.NewEdge(simple.Node(fromID), simple.Node(toID)))
	}

	return g
}

func parseInt(s string) int {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	if err != nil {
		panic(err)
	}
	return i
}

// Compare the Dijkstra implementation with gonum's implementation
// starting from Vertex 1
// using the web-Google.txt file
func TestDijkstraCorrectnessWithGonumAsComperrasion(t *testing.T) {
	const START = "1"
	// Load the web graph into both implementations
	gonumGraph := loadGonumGraph("web-Google.txt")
	ourGraph := graph.InitWebgraph()

	// Run Dijkstra's algorithm on both implementations starting from vertex "1"
	ourDistances := ourGraph.Dijkstra(START)
	if ourDistances == nil {
		t.Fatal("Our implementation failed to run Dijkstra's algorithm - start node may not exist")
	}

	// Run Gonum's Dijkstra implementation
	gonumDijkstra := path.DijkstraFrom(simple.Node(parseInt(START)), gonumGraph)

	// Checking if the same amount of nodes are reachable
	// Count reachable nodes in our implementation
	ourCount := 0
	for _, dist := range ourDistances {
		if !math.IsInf(dist, 1) {
			ourCount++
		}
	}
	gonumNodes := gonumGraph.Nodes()
	gonumCount := 0
	for gonumNodes.Next() {
		node := gonumNodes.Node().(simple.Node)
		if !math.IsInf(gonumDijkstra.WeightTo(int64(node)), 1) {
			gonumCount++
		}
	}
	if ourCount != gonumCount {
		t.Errorf("Our implementation found %d reachable nodes, but Gonum found %d", ourCount, gonumCount)
	}

	// Compare the results
	for node := range ourDistances {
		// Convert our node ID to int64 for Gonum
		nodeID, err := strconv.ParseInt(node, 10, 64)
		if err != nil {
			t.Errorf("Failed to parse node ID %s: %v", node, err)
			continue
		}

		// Get distance from Gonum
		gonumDist := gonumDijkstra.WeightTo(nodeID)

		// If Gonum can't reach the node but we can, or vice versa, that's an error
		if math.IsInf(gonumDist, 1) && !math.IsInf(ourDistances[node], 1) {
			t.Errorf("Gonum can't reach node %s but our implementation can", node)
		}
		if !math.IsInf(gonumDist, 1) && math.IsInf(ourDistances[node], 1) {
			t.Errorf("Our implementation can't reach node %s but Gonum can", node)
		}

		// If both can reach the node, compare distances
		if !math.IsInf(gonumDist, 1) && !math.IsInf(ourDistances[node], 1) {
			// Allow for small floating point differences
			if math.Abs(gonumDist-ourDistances[node]) > 1e-10 {
				t.Errorf("Distance mismatch for node %s: our implementation=%v, Gonum=%v",
					node, ourDistances[node], gonumDist)
				break
			}
		}
	}

	// Also check nodes that Gonum can reach but our implementation might not
	gonumNodes = gonumGraph.Nodes()
	for gonumNodes.Next() {
		node := gonumNodes.Node().(simple.Node)
		nodeStr := strconv.FormatInt(int64(node), 10)
		if _, exists := ourDistances[nodeStr]; !exists {
			t.Errorf("Our implementation didn't find node %s that exists in Gonum's graph", nodeStr)
		}
	}
}
