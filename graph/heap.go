package graph

type HeapNode struct {
	NodeId   string
	Distance float64
}

type minHeap struct {
	nodes []*HeapNode
}

// erstellt einen MinHeap und gibt diesen zurück
func newMinHeap(size int) *minHeap {
	return &minHeap{
		nodes: make([]*HeapNode, 0, size),
	}
}

func (h *minHeap) enqueue(node *HeapNode) {
	h.nodes = append(h.nodes, node)
	h.shiftUp(len(h.nodes) - 1)
}

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

func (h *minHeap) shiftUp(index int) {
	parentIndex := (index - 1) / 2

	for index > 0 && h.nodes[index].Distance < h.nodes[parentIndex].Distance {
		h.nodes[index], h.nodes[parentIndex] = h.nodes[parentIndex], h.nodes[index]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

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
