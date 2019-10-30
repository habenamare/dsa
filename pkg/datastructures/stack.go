package datastructures

import (
	"errors"
)

type stackItem struct {
	data interface{}
	next *stackItem
}

// A Stack data structure
type Stack struct {
	top  *stackItem
	size int
	max  int
}

// NewStack returns a Stack with unlimited capacity.
func NewStack() *Stack {
	return &Stack{}
}

// NewLimitedStack returns a Stack with limited capacity.
func NewLimitedStack(capacity int) *Stack {
	if capacity <= 0 {
		panic(errors.New("invalid capacity: capacity can not be zero or negative"))
	}

	return &Stack{max: capacity}
}

// IsEmpty reports whether the Stack has no items in it.
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Size returns the number of items currently found inside the stack.
func (s *Stack) Size() int {
	return s.size
}

// isLimited reports whether the Stack has limited capacity.
func (s *Stack) isLimited() bool {
	return s.max != 0
}

// IsFull reports whether items can be pushed into the Stack.
func (s *Stack) IsFull() bool {
	if !s.isLimited() {
		return false
	}

	return s.size == s.max
}

// Push adds a new item to the top of the Stack.
func (s *Stack) Push(data interface{}) {
	if s.IsFull() {
		panic(errors.New("stack is full"))
	}

	newTop := &stackItem{
		data: data,
		next: s.top,
	}

	s.top = newTop
	s.size++
}

// Peek returns the item found on the top of the Stack, i.e. the
// last pushed item.
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}

	return s.top.data
}

// Pop removes and returns the item found on the top of the Stack.
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}

	topData := s.top.data

	s.top = s.top.next
	s.size--

	return topData
}
