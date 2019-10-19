package algorithms

import (
	"testing"
)

func TestBracketsMatch(t *testing.T) {
	for _, tc := range []struct {
		brackets string
		result   bool
	}{
		{brackets: "{}", result: true},
		{brackets: "[{}{}]", result: true},
		{brackets: "{)", result: false},
		{brackets: "{[()]}))[]", result: false},
		{brackets: "{}[]{([])}", result: true},
	} {
		if got := BracketsMatch(tc.brackets); got != tc.result {
			t.Fatalf(
				"expected the matching of brackets in %s to be %t, got %t\n",
				tc.brackets,
				tc.result,
				got,
			)
		}
	}
}

func TestBracketsMatch_panic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			if err.Error() != "invalid input" {
				t.Fatal("unexpected panic")
			}

		} else {
			t.Fatal("the function did not panic")
		}
	}()

	BracketsMatch("{x")
}
