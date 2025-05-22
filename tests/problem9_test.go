package tests

import (
	"AbgabeAlgo/graph"
	"testing"
)

func TestProblem9test(t *testing.T) {
	problem9 := graph.InitProblem98test()

	distances := problem9.Dijkstra("1")

	expectedDistances := make(map[string]float64)
	expectedDistances["1"] = 0.0
	expectedDistances["2"] = 1.0
	expectedDistances["3"] = 2.0
	expectedDistances["4"] = 3.0
	expectedDistances["5"] = 4.0
	expectedDistances["6"] = 4.0
	expectedDistances["7"] = 3.0
	expectedDistances["8"] = 2.0

	if len(distances) == len(expectedDistances) {
		for id, d := range distances {
			if expectedDistances[id] != d {
				t.Errorf("Dijkstra method error: expected distance %v for node %v, got %v", expectedDistances[id], id, d)
			}
		}
	} else {
		t.Errorf("Dijkstra method error: expected %d distances, got %d", len(expectedDistances), len(distances))
	}

}
