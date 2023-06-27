package graph

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// ParseGraphFromFile parses the given file to create a DirectedGraph.
func ParseGraphFromFile(filename string) *DirectedGraph {
	file, _ := os.Open(filename)

	defer file.Close()

	graph := &DirectedGraph{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 1 {
			continue
		}

		from, _ := strconv.Atoi(fields[0])

		fromKey := fmt.Sprintf("%d", from)
		graph.AddVertex(fromKey)

		for _, field := range fields[1:] {
			parts := strings.Split(field, ",")
			if len(parts) != 2 {
				continue
			}

			to, _ := strconv.Atoi(parts[0])

			length, _ := strconv.ParseFloat(parts[1], 64)

			toKey := fmt.Sprintf("%d", to)
			graph.AddVertex(toKey)
			graph.AddDirectedEdge(fromKey, toKey, length)
		}
	}

	return graph
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
