package graph

type Stack []*Vertex

func (s *Stack) Push(v *Vertex) {
	*s = append(*s, v)
}

func (stack *Stack) Pop() *Vertex {
	res := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return res
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
