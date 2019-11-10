package datastructures

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()

	// Assert that `s` points to an instance of a `Stack`
	if _, ok := interface{}(*s).(Stack); !ok {
		t.Fatalf(
			"expected `s` obtained from `NewStack()` to point to an instance of a %v, "+
				"instead it points to an instance of a %v",
			"`datastructures.Stack`",
			fmt.Sprintf("`%s`", reflect.TypeOf(*s)),
		)
	}

	if s.top != nil {
		t.Fatalf(
			"expected `top` of Stack obtained from `NewStack()` to be %v, got %v",
			nil,
			s.top,
		)
	}

	if s.size != 0 {
		t.Fatalf(
			"expected `size` of Stack obtained from `NewStack()` to be %d, got %d",
			0,
			s.size,
		)
	}

	if s.max != 0 {
		t.Fatalf(
			"expected `max` of Stack obtained from `NewStack()` to be %d, got %d",
			0,
			s.max,
		)
	}
}

func TestNewLimitedStack(t *testing.T) {
	limitedS := NewLimitedStack(5)

	// Assert that `limitedS` points to an instance of a `Stack`
	if _, ok := interface{}(*limitedS).(Stack); !ok {
		t.Fatalf(
			"expected `limitedS` obtained from `NewLimitedStack(5)` to point to an instance of a %v, "+
				"instead it points to an instance of a %v",
			"`datastructures.Stack`",
			fmt.Sprintf("`%s`", reflect.TypeOf(*limitedS)),
		)
	}

	if limitedS.top != nil {
		t.Fatalf(
			"expected `top` of Stack obtained from `NewLimitedStack(5)` to be %v, got %v",
			nil,
			limitedS.top,
		)
	}

	if limitedS.size != 0 {
		t.Fatalf(
			"expected `size` of Stack obtained from `NewLimitedStack(5)` to be %d, got %d",
			0,
			limitedS.size,
		)
	}

	if limitedS.max != 5 {
		t.Fatalf(
			"expected `max` of Stack obtained from `NewLimitedStack(5)` to be %d, got %d",
			5,
			limitedS.max,
		)
	}
}

func TestIsEmpty(t *testing.T) {
	s := &Stack{size: 0}

	if !s.IsEmpty() {
		t.Fatalf(
			"expected calling `isEmpty()` on a Stack with a `size` of 0 to return %v, got %v",
			true,
			s.IsEmpty(),
		)
	}
}

func TestSize(t *testing.T) {
	s := &Stack{size: 5}

	if s.Size() != 5 {
		t.Fatalf(
			"expected calling `Size()` on a Stack with a `size` of 5 to return %v, got %v",
			5,
			s.Size(),
		)
	}
}

func TestIsFull(t *testing.T) {
	testCases := []struct {
		s        Stack
		expected bool
	}{
		{Stack{max: 0, size: 100}, false},
		{Stack{max: 5, size: 4}, false},
		{Stack{max: 5, size: 5}, true},
	}

	for _, tc := range testCases {
		t.Run(
			fmt.Sprintf("Stack{max: %d, size: %d}", tc.s.max, tc.s.size),
			func(t *testing.T) {
				if got := tc.s.IsFull(); got != tc.expected {
					t.Fatalf(
						"expected calling `IsFull()` on a Stack with a `max` of %v and a `size` of %v to return %v, got %v",
						tc.s.max,
						tc.s.size,
						tc.expected,
						tc.s.IsFull(),
					)
				}
			},
		)
	}
}

func TestPush(t *testing.T) {
	s := &Stack{top: nil, max: 0, size: 0}

	s.Push("first item")
	if s.top.data != "first item" {
		t.Fatalf(
			"expected calling `Push(\"first item\")` on an empty Stack to make `%v` the data on top of the Stack, got `%v` on top of the Stack instead",
			"first item",
			s.top.data,
		)
	}

	if s.size != 1 {
		t.Fatalf(
			"expected `size` of Stack to increase from 0 to %v after calling `Push(\"first item\")`, instead the `size` is now %v",
			1,
			s.size,
		)
	}
}

func TestPush_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf(
				"expected calling `Push(\"item\")` on a full Stack to cause a panic",
			)
		}
	}()

	fullS := &Stack{max: 3, size: 3}
	fullS.Push("item")
}

func TestPeek(t *testing.T) {
	s := &Stack{top: &stackItem{data: "top item"}, size: 1}

	topItem := s.Peek()
	if topItem == nil {
		t.Fatalf(
			"expected calling `Peek()` on a Stack with a top item to return `%v`, got `%v`",
			"top item",
			topItem,
		)
	}

	if s.size != 1 {
		t.Fatalf("expected no change to the `size` of Stack after a call to `Peek()`")
	}
}

func TestPeek_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf(
				"expected calling `Peek()` on a Stack with a `size` of 0 to cause a panic",
			)
		}
	}()

	s := &Stack{size: 0}
	s.Peek()
}

func TestPop(t *testing.T) {
	s := &Stack{top: &stackItem{data: "top item"}, size: 1}

	topItem := s.Pop()
	if topItem == nil {
		t.Fatalf(
			"expected calling `Pop()` on a Stack with a top item to return `%v`, got `%v`",
			"top item",
			topItem,
		)
	}

	if s.size != 0 {
		t.Fatalf(
			"expected `size` of Stack to decrease from 1 to %v after a call to `Pop()`, instead the `size` is now %v",
			0,
			s.size,
		)
	}
}

func TestPop_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf(
				"expected calling `Pop()` on a Stack with a `size` of 0 to cause a panic",
			)
		}
	}()

	s := &Stack{size: 0}
	s.Pop()
}
