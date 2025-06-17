package graph

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// initGraph9 initializes an undirected graph from a file.
// Format: vertex_id edge1_id,weight1 edge2_id,weight2 ...
// Returns a pointer to the initialized UnDirectedGraph.
func initGraph9(filename string) *UnDirectedGraph {
	graph := &UnDirectedGraph{}

	totalLines, err := countLines(filename)
	if err != nil {
		panic(fmt.Sprintf("Can't count lines in file %s: %v", filename, err))
	}

	file, ok := os.Open(filename)
	if ok != nil {
		errorString := fmt.Sprintf("Can't open file %s", filename)
		panic(errorString)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validLines := 0
	nextPercent := 5
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		if len(fields) == 0 {
			continue
		}
		id1 := fields[0]
		graph.AddVertex(id1)
		for _, x := range fields[1:] {
			f := strings.Split(x, ",")
			var length float64
			if l, err := strconv.ParseFloat(f[1], 64); err == nil {
				length = l
			} else {
				panic("convert str2float failed!")
			}
			graph.AddVertex(f[0])
			graph.AddUndirectedEdge(id1, f[0], length)
		}

		validLines++
		// Print percentage progress every 5% or every 50,000 valid lines
		percentDone := int(float64(validLines) / float64(totalLines) * 100)
		if percentDone >= nextPercent || validLines%50000 == 0 {
			fmt.Printf("Progress: %d%% (%d/%d lines)\n", percentDone, validLines, totalLines)
			if percentDone >= nextPercent {
				nextPercent += 5
			}
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

	// This is to ensure the path is correct when running tests
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
// The function includes progress tracking and time estimation.
// Returns a pointer to the initialized DirectedGraph.
func InitWebgraph() *DirectedGraph {
	webgraph := DirectedGraph{}

	// Get the absolute path to the testdata directory
	wd, err := os.Getwd()
	if err != nil {
		panic("Could not get working directory")
	}
	// This is to ensure the path is correct when running tests
	if filepath.Base(wd) == "tests" {
		wd = filepath.Dir(wd)
	}
	filepath := filepath.Join(wd, "testdata", "web-Google.txt")

	totalLines, err := countLines(filepath)
	if err != nil {
		panic(fmt.Sprintf("Could not count lines in file %s: %v", filepath, err))
	}

	file, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %s: %v", filepath, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validLines := 0
	nextPercent := 5
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

		validLines++

		webgraph.AddVertex(fields[0])
		webgraph.AddVertex(fields[1])
		webgraph.AddDirectedEdge(fields[0], fields[1], 1.)

		// Print percentage progress every 5% or every 50,000 valid lines
		percentDone := int(float64(validLines) / float64(totalLines) * 100)
		if percentDone >= nextPercent || validLines%50000 == 0 {
			fmt.Printf("Progress: %d%% (%d/%d lines)\n", percentDone, validLines, totalLines)
			if percentDone >= nextPercent {
				nextPercent += 5
			}
		}
	}

	return &webgraph
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
