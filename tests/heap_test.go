package tests

import (
	"AbgabeAlgo/graph"
	"testing"
)

// TestNewMinHeap tests the creation of a new min heap
func TestNewMinHeap(t *testing.T) {
	heap := graph.NewMinHeap(10)
	if heap == nil {
		t.Error("newMinHeap returned nil")
	}
	if !heap.IsEmpty() {
		t.Error("New heap should be empty")
	}
}

// TestEnqueueDequeue tests the enqueue and dequeue operations of a min heap
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

// TestContains tests the contains operation of a min heap
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

func TestUpdate(t *testing.T) {
	heap := graph.NewMinHeap(4)
	nodes := []*graph.HeapNode{
		{NodeId: "A", Distance: 10.0},
		{NodeId: "B", Distance: 20.0},
		{NodeId: "C", Distance: 30.0},
		{NodeId: "D", Distance: 40.0},
	}
	for _, node := range nodes {
		heap.Enqueue(node)
	}

	// Decrease key: C from 30.0 to 5.0 (should move to root)
	err := heap.Update("C", 5.0)
	if err != nil {
		t.Errorf("Unexpected error when decreasing key: %v", err)
	}
	node := heap.Dequeue()
	if node.NodeId != "C" {
		t.Errorf("Expected C to be at root after decrease, got %s", node.NodeId)
	}

	// Increase key: A from 10.0 to 50.0 (should move to bottom)
	err = heap.Update("A", 50.0)
	if err != nil {
		t.Errorf("Unexpected error when increasing key: %v", err)
	}
	// B should now be at root
	node = heap.Dequeue()
	if node.NodeId != "B" {
		t.Errorf("Expected B to be at root after increasing A, got %s", node.NodeId)
	}

	// Update with same value: D from 40.0 to 40.0 (should not change position)
	err = heap.Update("D", 40.0)
	if err != nil {
		t.Errorf("Unexpected error when updating with same value: %v", err)
	}

	// Update non-existent node
	err = heap.Update("X", 100.0)
	if err == nil {
		t.Error("Expected error when updating non-existent node, got nil")
	}
}
