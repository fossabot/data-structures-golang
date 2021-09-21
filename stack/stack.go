// Stack is a linear data structure which follows a particular order in which
// the operations are performed. The order may be LIFO(Last In First Out) or
// FILO(First In Last Out).
package stack

// Stack represents a queue that holds a slice
type Stack struct {
	Items []int
}

// Push will add value at the end
func (s *Stack) Push(i int) {
	s.Items = append(s.Items, i)
}

// Pop will remove a value at the end and return the removed value
func (s *Stack) Pop() int {
	l := len(s.Items) - 1
	toRemove := s.Items[l]
	s.Items = s.Items[:l]
	return toRemove
}

// Create a new Stack
func NewStack() Stack {
	return Stack{}
}
