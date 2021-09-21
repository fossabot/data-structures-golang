// Stack is a linear data structure which follows a particular order in which
// the operations are performed. The order may be LIFO(Last In First Out) or
// FILO(First In Last Out).
package stack

// Stack represents a queue that holds a slice
type Stack struct {
	items []int
}

// Push will add value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end and return the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

// Create a new Stack
func NewStack() Stack {
	return Stack{}
}
