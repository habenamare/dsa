package datastructures

import (
	"errors"
)

type stackItem struct {
	data interface{}
	next *stackItem
}

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

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsFull() bool {
	if s.max == 0 {
		return false
	}

	return s.size == s.max
}

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

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}

	topData := s.top.data

	s.top = s.top.next
	s.size--

	return topData
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		panic(errors.New("stack is empty"))
	}

	return s.top.data
}
