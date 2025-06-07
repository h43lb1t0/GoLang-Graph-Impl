package graph

import "fmt"

// HeapNode represents a node in the min heap with an associated distance value.
// It is used for implementing Dijkstra's algorithm where NodeId identifies the vertex
// and Distance represents the current shortest distance to that vertex.
type HeapNode struct {
	NodeId   string
	Distance float64
}

// MinHeap implements a binary min heap data structure.
// It now includes a map to track node positions for efficient decreaseKey operations.
type MinHeap struct {
	nodes     []*HeapNode
	positions map[string]int // Maps NodeId to its index in the 'nodes' slice
}

// NewMinHeap creates and returns a new min heap with the specified initial capacity.
func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		nodes:     make([]*HeapNode, 0, size),
		positions: make(map[string]int),
	}
}

// IsEmpty checks if the heap contains no elements.
func (h *MinHeap) IsEmpty() bool {
	return len(h.nodes) == 0
}

// swap is a helper function to swap two nodes in the heap
// and update their recorded positions.
func (h *MinHeap) swap(i, j int) {
	// Swap positions in the map
	h.positions[h.nodes[i].NodeId] = j
	h.positions[h.nodes[j].NodeId] = i
	// Swap nodes in the slice
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// Enqueue adds a new node to the heap and maintains the heap property.
// It also records the node's position.
func (h *MinHeap) Enqueue(node *HeapNode) {
	h.nodes = append(h.nodes, node)
	index := len(h.nodes) - 1
	h.positions[node.NodeId] = index // Store initial position
	h.shiftUp(index)
}

// Dequeue removes and returns the node with the minimum distance from the heap.
// Returns nil if the heap is empty. It updates positions accordingly.
func (h *MinHeap) Dequeue() *HeapNode {
	if h.IsEmpty() {
		return nil
	}

	nodeToReturn := h.nodes[0]
	delete(h.positions, nodeToReturn.NodeId) // Node is no longer in heap

	numNodes := len(h.nodes)
	if numNodes == 1 {
		h.nodes = h.nodes[:0] // Empty the slice
		return nodeToReturn
	}

	// Move the last node to the root
	h.nodes[0] = h.nodes[numNodes-1]
	h.positions[h.nodes[0].NodeId] = 0 // Update position of the node now at root

	h.nodes = h.nodes[:numNodes-1] // Shrink slice

	h.shiftDown(0) // Restore heap property from the root

	return nodeToReturn
}

// shiftUp maintains the heap property by moving a node up the heap.
func (h *MinHeap) shiftUp(index int) {
	parentIndex := (index - 1) / 2
	// Loop while the current node is not the root and is smaller than its parent
	for index > 0 && h.nodes[index].Distance < h.nodes[parentIndex].Distance {
		h.swap(index, parentIndex)
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

// shiftDown maintains the heap property by moving a node down the heap.
func (h *MinHeap) shiftDown(index int) {
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild < len(h.nodes) && h.nodes[leftChild].Distance < h.nodes[smallest].Distance {
			smallest = leftChild
		}
		if rightChild < len(h.nodes) && h.nodes[rightChild].Distance < h.nodes[smallest].Distance {
			smallest = rightChild
		}

		if smallest == index { // The node is in the correct position
			break
		}

		h.swap(index, smallest)
		index = smallest // Move down to the swapped child's position
	}
}

// Contains checks if a node with the given NodeId is currently in the heap.
func (h *MinHeap) Contains(nodeId string) bool {
	_, exists := h.positions[nodeId]
	return exists
}

// DecreaseKey updates the distance of a node already in the heap and adjusts its position.
// It's crucial that newDistance is actually smaller than the current distance.
func (h *MinHeap) DecreaseKey(nodeId string, newDistance float64) error {
	index, exists := h.positions[nodeId]
	if !exists {
		return fmt.Errorf("node %s not found in heap for decreaseKey", nodeId)
	}

	// This check ensures we are actually "decreasing" the key.
	// If Dijkstra's algorithm is correct, this condition (newDistance >= current)
	// should ideally not be met when decreaseKey is called.
	if newDistance >= h.nodes[index].Distance {
		return fmt.Errorf("new distance %.2f for node %s is not smaller than current distance %.2f",
			newDistance, nodeId, h.nodes[index].Distance)
	}

	h.nodes[index].Distance = newDistance
	h.shiftUp(index) // After decreasing a key, the node might need to move up.
	return nil
}
