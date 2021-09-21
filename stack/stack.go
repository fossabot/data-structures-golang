package stack

type Stack struct {
	items []int
}

// Push will add value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// and RETURNS the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func NewStack() Stack {
	return Stack{}
}
