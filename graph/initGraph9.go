package graph

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

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

func InitProblem98() *UnDirectedGraph {
	wd, err := os.Getwd()
	if err != nil {
		panic("Could not get working directory")
	}
	return initGraph9(filepath.Join(wd, "testdata", "problem9.8.txt"))
}

func InitWebgraph() *DirectedGraph {
	// for a sanity check:
	//     count on the command line the number of edges and vertices by
	// grep -E -v "^#" ~/Downloads/web-Google.txt | wc -l
	// grep -E -v "^#" ~/Downloads/web-Google.txt | sed -E 's/([[:digit:]]+)[[:space:]]+([[:digit:]]+).*/\1\n\2/' | sort | uniq | wc -l

	webgraph := DirectedGraph{}

	file, _ := os.Open("./testdata/web-Google.txt")
	defer file.Close()

	fmt.Println("reading file")

	scanner := bufio.NewScanner(file)
	i := 1
	start := time.Now()
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		if !strings.HasPrefix((fields[0][0:1]), "#") && len(fields) == 2 {
			webgraph.AddVertex(fields[0])
			webgraph.AddVertex(fields[1])
			webgraph.AddDirectedEdge(fields[0], fields[1], 1.)
		}
		if (i % 1000000) == 0 {
			elapsed := time.Since(start)
			fmt.Printf("last took %s\n", elapsed)
			fmt.Printf("progess: %v\n", i)
			start = time.Now()
		}
		if (i % 50000) == 0 {
			printEstimatedRemainingTime(start, i)
		}
		i++
	}
	fmt.Printf("%v lines processed\n", i)

	return &webgraph
}

func printEstimatedRemainingTime(start time.Time, numProcessed int) {
	total := -1
	if total == -1 {
		total, _ = countLines("./testdata/web-Google.txt")
	}
	elapsed := time.Since(start)
	rate := float64(numProcessed) / elapsed.Seconds()   // processing rate in lines per second
	remaining := total - numProcessed                   // remaining lines to process
	estimatedTimeRemaining := float64(remaining) / rate // estimated remaining time in seconds
	fmt.Printf("Estimated remaining time: %s\n", time.Duration(estimatedTimeRemaining)*time.Second)
}

func countLines(filepath string) (int, error) {
	out, err := exec.Command("wc", "-l", filepath).Output()
	if err != nil {
		return 0, err
	}

	// The output will be something like "200 filename", so we need to parse out the number
	split := strings.Split(string(out), " ")
	lines, err := strconv.Atoi(split[0])
	fmt.Printf("lines: %v\n", lines)
	return lines, err
}
