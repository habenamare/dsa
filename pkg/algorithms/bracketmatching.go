package algorithms

import (
	"errors"
	"fmt"

	"github.com/habenamare/dsa/pkg/datastructures"
)

func BracketsMatch(brackets string) bool {
	stack := datastructures.NewStack()

	for i := 0; i < len(brackets); i++ {
		bracketChar := rune(brackets[i])

		if isLeftBracket(bracketChar) {
			stack.Push(bracketChar)
		} else if isRightBracket(bracketChar) {
			if stack.IsEmpty() {
				return false
			}

			reversedBracket := getReversedBracket(bracketChar)

			stackTop := stack.Pop()
			if stackTop != reversedBracket {
				return false
			}
		} else {
			panic(errors.New("invalid input"))
		}
	}

	return stack.IsEmpty()
}

func isLeftBracket(bracket rune) bool {
	if bracket == '(' || bracket == '[' || bracket == '{' {
		return true
	}

	return false
}

func isRightBracket(bracket rune) bool {
	if bracket == ')' || bracket == ']' || bracket == '}' {
		return true
	}

	return false
}

func getReversedBracket(bracket rune) rune {
	switch bracket {
	case '(':
		return ')'
	case ')':
		return '('
	case '[':
		return ']'
	case ']':
		return '['
	case '{':
		return '}'
	case '}':
		return '{'
	}

	panic(fmt.Sprintf(" %c  is not a bracket", bracket))
}
