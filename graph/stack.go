package graph

// Stack represents a LIFO (Last In First Out) data structure for storing vertices.
type Stack []*Vertex

// Push adds a vertex to the top of the stack.
func (s *Stack) Push(v *Vertex) {
	*s = append(*s, v)
}

// Pop removes and returns the vertex from the top of the stack.
// It assumes the stack is not empty.
func (stack *Stack) Pop() *Vertex {
	res := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return res
}

// IsEmpty returns true if the stack has no elements.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
