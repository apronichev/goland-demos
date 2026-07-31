package stack

// Stack is a generic LIFO stack.
type Stack[T any] struct {
	items []T
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// TODO: Implement support for converting all values in the stack using a caller-provided function.
