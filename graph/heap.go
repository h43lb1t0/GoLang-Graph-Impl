package graph

// HeapNode represents a node in the min heap with an associated distance value.
// It is used for implementing Dijkstra's algorithm where NodeId identifies the vertex
// and Distance represents the current shortest distance to that vertex.
type HeapNode struct {
	NodeId   string
	Distance float64
}

// minHeap implements a binary min heap data structure.
// The heap maintains the property that each parent node has a smaller distance
// than its children, making it efficient for retrieving the minimum element.
type minHeap struct {
	nodes []*HeapNode
}

// newMinHeap creates and returns a new min heap with the specified initial capacity.
// The heap starts empty but can grow up to the specified size without reallocation.
func newMinHeap(size int) *minHeap {
	return &minHeap{
		nodes: make([]*HeapNode, 0, size),
	}
}

// enqueue adds a new node to the heap and maintains the heap property
// by shifting the node up to its correct position.
func (h *minHeap) enqueue(node *HeapNode) {
	h.nodes = append(h.nodes, node)
	h.shiftUp(len(h.nodes) - 1)
}

// dequeue removes and returns the node with the minimum distance from the heap.
// Returns nil if the heap is empty.
// After removal, it maintains the heap property by shifting down the last element.
func (h *minHeap) dequeue() *HeapNode {
	if len(h.nodes) == 0 {
		return nil
	}

	node := h.nodes[0]
	lastNode := h.nodes[len(h.nodes)-1]
	h.nodes[0] = lastNode
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.shiftDown(0)

	return node
}

// shiftUp maintains the heap property by moving a node up the heap
// until its parent has a smaller distance value.
func (h *minHeap) shiftUp(index int) {
	parentIndex := (index - 1) / 2

	for index > 0 && h.nodes[index].Distance < h.nodes[parentIndex].Distance {
		h.nodes[index], h.nodes[parentIndex] = h.nodes[parentIndex], h.nodes[index]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

// shiftDown maintains the heap property by moving a node down the heap
// until it is smaller than both its children or becomes a leaf node.
func (h *minHeap) shiftDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	smallest := index
	if leftChild < len(h.nodes) && h.nodes[leftChild].Distance < h.nodes[smallest].Distance {
		smallest = leftChild
	}
	if rightChild < len(h.nodes) && h.nodes[rightChild].Distance < h.nodes[smallest].Distance {
		smallest = rightChild
	}
	if smallest != index {
		h.nodes[index], h.nodes[smallest] = h.nodes[smallest], h.nodes[index]
		h.shiftDown(smallest)
	}
}
