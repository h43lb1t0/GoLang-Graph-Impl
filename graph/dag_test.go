package graph

import (
	"reflect"
	"testing"
)

func TestAddDirectedEdgeDAG(t *testing.T) {
	dag := DAG{}
	dag.AddVertex("1")
	dag.AddVertex("2")
	dag.AddDirectedEdge("1", "2", 1.0)

	if len(dag.vertices[0].adjacent) != 1 {
		t.Errorf("Failed to add a directed edge")
	}
}

func TestTopoSort(t *testing.T) {
	dag := DAG{}
	dag.AddVertex("1")
	dag.AddVertex("2")
	dag.AddVertex("3")
	dag.AddDirectedEdge("1", "2", 1.0)
	dag.AddDirectedEdge("2", "3", 1.0)
	ordering := dag.TopoSort()

	expected := map[string]int{
		"1": 3,
		"2": 2,
		"3": 1,
	}

	if !reflect.DeepEqual(ordering, expected) {
		t.Errorf("TopoSort method error, got: %v, want: %v.", ordering, expected)
	}
}

func TestTopoSortNotDag(t *testing.T) {
	dag := DAG{}
	dag.AddVertex("1")
	dag.AddVertex("2")
	dag.AddVertex("3")
	dag.AddDirectedEdge("1", "2", 1.0)
	dag.AddDirectedEdge("2", "3", 1.0)
	dag.AddDirectedEdge("3", "1", 1.0)
	ordering := dag.TopoSort()

	expected := map[string]int{
		"1": 3,
		"2": 2,
		"3": 1,
	}

	if reflect.DeepEqual(ordering, expected) {
		t.Errorf("TopoSort method error, got: %v, want: %v.", ordering, expected)
	}
}
