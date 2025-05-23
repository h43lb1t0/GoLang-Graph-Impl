package graph

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// initGraph9 initializes an undirected graph from a file with a specific format.
// The file should contain lines where each line represents a vertex and its edges.
// Format: vertex_id edge1_id,weight1 edge2_id,weight2 ...
// Returns a pointer to the initialized UnDirectedGraph.
func initGraph9(filename string) *UnDirectedGraph {
	graph := &UnDirectedGraph{}

	file, ok := os.Open(filename)
	if ok != nil {
		errorString := fmt.Sprintf("Can't open file %s", filename)
		panic(errorString)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		id1 := fields[0]
		graph.AddVertex(id1)
		for _, x := range fields[1:] {
			f := strings.Split(x, ",") // f[0]:id2 ,,
			var length float64         //edgeLength
			if l, err := strconv.ParseFloat(f[1], 64); err == nil {
				length = l //edgeLength{float64: l}
			} else {
				panic("convert str2float failed!")
			}
			graph.AddVertex(f[0])
			graph.AddUndirectedEdge(id1, f[0], length)
		}
	}
	return graph
}

// InitProblem98test initializes a test graph for problem 9.8.
// It reads the graph data from "testdata/problem9.8test.txt".
// Returns a pointer to the initialized UnDirectedGraph.
func InitProblem98test() *UnDirectedGraph {
	wd, err := os.Getwd()
	if err != nil {
		panic("Could not get working directory")
	}
	// If we're in the tests directory, go up one level
	if filepath.Base(wd) == "tests" {
		wd = filepath.Dir(wd)
	}
	return initGraph9(filepath.Join(wd, "testdata", "problem9.8test.txt"))
}

// InitProblem98 initializes the main graph for problem 9.8.
// It reads the graph data from "testdata/problem9.8.txt".
// Returns a pointer to the initialized UnDirectedGraph.
func InitProblem98() *UnDirectedGraph {
	wd, err := os.Getwd()
	if err != nil {
		panic("Could not get working directory")
	}
	return initGraph9(filepath.Join(wd, "testdata", "problem9.8.txt"))
}

// InitWebgraph initializes a directed graph from the Google web graph dataset.
// The dataset should be in the format of "testdata/web-Google.txt".
// Each line represents a directed edge from source to target vertex.
// The function includes progress tracking and time estimation.
// Returns a pointer to the initialized DirectedGraph.
func InitWebgraph() *DirectedGraph {
	webgraph := DirectedGraph{}

	// Get the absolute path to the testdata directory
	wd, err := os.Getwd()
	if err != nil {
		panic("Could not get working directory")
	}
	// If we're in the tests directory, go up one level
	if filepath.Base(wd) == "tests" {
		wd = filepath.Dir(wd)
	}
	filepath := filepath.Join(wd, "testdata", "web-Google.txt")

	// Count total lines first
	totalLines, err := countLines(filepath)
	if err != nil {
		panic(fmt.Sprintf("Could not count lines in file %s: %v", filepath, err))
	}

	file, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %s: %v", filepath, err))
	}
	defer file.Close()

	fmt.Println("reading file")

	scanner := bufio.NewScanner(file)
	i := 1
	start := time.Now()
	for scanner.Scan() {
		s := scanner.Text()
		// Skip comment lines and empty lines
		if strings.HasPrefix(s, "#") || strings.TrimSpace(s) == "" {
			continue
		}
		fields := strings.Fields(s)
		if len(fields) != 2 {
			continue
		}

		// Only add vertices if they don't exist
		if !webgraph.containsVertex(webgraph.vertices, fields[0]) {
			webgraph.AddVertex(fields[0])
		}
		if !webgraph.containsVertex(webgraph.vertices, fields[1]) {
			webgraph.AddVertex(fields[1])
		}
		webgraph.AddDirectedEdge(fields[0], fields[1], 1.)

		if (i % 1000000) == 0 {
			elapsed := time.Since(start)
			fmt.Printf("last took %s\n", elapsed)
			fmt.Printf("progess: %v\n", i)
			start = time.Now()
		}
		if (i % 50000) == 0 {
			printEstimatedRemainingTime(start, i, totalLines)
		}
		i++
	}
	fmt.Printf("%v lines processed\n", i)

	return &webgraph
}

// printEstimatedRemainingTime calculates and prints the estimated time remaining
// for processing the web graph based on the current processing rate.
// Parameters:
//   - start: The start time of the processing
//   - numProcessed: The number of lines processed so far
//   - totalLines: The total number of lines to process
func printEstimatedRemainingTime(start time.Time, numProcessed int, totalLines int) {
	elapsed := time.Since(start)
	rate := float64(numProcessed) / elapsed.Seconds()   // processing rate in lines per second
	remaining := totalLines - numProcessed              // remaining lines to process
	estimatedTimeRemaining := float64(remaining) / rate // estimated remaining time in seconds
	fmt.Printf("Estimated remaining time: %s\n", time.Duration(estimatedTimeRemaining)*time.Second)
}

// countLines counts the number of lines in a file.
// Parameters:
//   - filepath: The path to the file to count lines in
//
// Returns:
//   - int: The number of lines in the file
//   - error: Any error that occurred during the operation
func countLines(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	fmt.Printf("Total lines: %v\n", lineCount)
	return lineCount, nil
}
