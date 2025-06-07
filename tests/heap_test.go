package tests

import (
	"AbgabeAlgo/graph"
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	heap := graph.NewMinHeap(10)
	if heap == nil {
		t.Error("newMinHeap returned nil")
	}
	if !heap.IsEmpty() {
		t.Error("New heap should be empty")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	heap := graph.NewMinHeap(5)

	// Test enqueue
	nodes := []*graph.HeapNode{
		{NodeId: "A", Distance: 5.0},
		{NodeId: "B", Distance: 3.0},
		{NodeId: "C", Distance: 7.0},
		{NodeId: "D", Distance: 1.0},
	}

	for _, node := range nodes {
		heap.Enqueue(node)
	}

	// Test dequeue order (should be in ascending order of distance)
	expectedOrder := []string{"D", "B", "A", "C"}
	for _, expectedId := range expectedOrder {
		node := heap.Dequeue()
		if node == nil {
			t.Error("Dequeue returned nil")
			continue
		}
		if node.NodeId != expectedId {
			t.Errorf("Expected node %s, got %s", expectedId, node.NodeId)
		}
	}

	// Test empty heap
	if !heap.IsEmpty() {
		t.Error("Heap should be empty after all dequeues")
	}
	if heap.Dequeue() != nil {
		t.Error("Dequeue on empty heap should return nil")
	}
}

func TestDecreaseKey(t *testing.T) {
	heap := graph.NewMinHeap(3)

	// Add nodes
	nodes := []*graph.HeapNode{
		{NodeId: "A", Distance: 5.0},
		{NodeId: "B", Distance: 3.0},
		{NodeId: "C", Distance: 7.0},
	}

	for _, node := range nodes {
		heap.Enqueue(node)
	}

	// Test valid decrease key
	err := heap.DecreaseKey("C", 2.0)
	if err != nil {
		t.Errorf("Unexpected error in decreaseKey: %v", err)
	}

	// Verify new order after decrease key
	expectedOrder := []string{"C", "B", "A"}
	for _, expectedId := range expectedOrder {
		node := heap.Dequeue()
		if node.NodeId != expectedId {
			t.Errorf("Expected node %s, got %s", expectedId, node.NodeId)
		}
	}

	// Test decrease key on non-existent node
	err = heap.DecreaseKey("X", 1.0)
	if err == nil {
		t.Error("Expected error when decreasing key of non-existent node")
	}

	// Test decrease key with larger value
	heap = graph.NewMinHeap(1)
	heap.Enqueue(&graph.HeapNode{NodeId: "A", Distance: 3.0})
	err = heap.DecreaseKey("A", 4.0)
	if err == nil {
		t.Error("Expected error when new distance is larger than current distance")
	}
}

func TestContains(t *testing.T) {
	heap := graph.NewMinHeap(2)

	// Add a node
	node := &graph.HeapNode{NodeId: "A", Distance: 5.0}
	heap.Enqueue(node)

	// Test contains for existing node
	if !heap.Contains("A") {
		t.Error("Contains should return true for existing node")
	}

	// Test contains for non-existing node
	if heap.Contains("B") {
		t.Error("Contains should return false for non-existing node")
	}

	// Test contains after dequeue
	heap.Dequeue()
	if heap.Contains("A") {
		t.Error("Contains should return false for dequeued node")
	}
}
