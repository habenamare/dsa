package patternmatching

import (
	"fmt"
	"testing"
)

func TestBruteForceMatch(t *testing.T) {
	testCases := []struct {
		pattern  string
		text     string
		expected Matches
	}{
		{"a", "ababababa", Matches{
			NewMatch(0, 0),
			NewMatch(2, 2),
			NewMatch(4, 4),
			NewMatch(6, 6),
			NewMatch(8, 8),
		}},
		{"b", "ababababa", Matches{
			NewMatch(1, 1),
			NewMatch(3, 3),
			NewMatch(5, 5),
			NewMatch(7, 7),
		}},
		{"aba", "ababababa", Matches{
			NewMatch(0, 2),
			NewMatch(2, 4),
			NewMatch(4, 6),
			NewMatch(6, 8),
		}},
		{"bab", "ababababa", Matches{
			NewMatch(1, 3),
			NewMatch(3, 5),
			NewMatch(5, 7),
		}},
		{"a", "abbbbbbba", Matches{
			NewMatch(0, 0),
			NewMatch(8, 8),
		}},
	}

	for _, tc := range testCases {
		t.Run(
			fmt.Sprintf("Pattern: %s", tc.pattern),
			func(t *testing.T) {
				if got := BruteForceMatch(tc.pattern, tc.text); !got.Equals(tc.expected) {
					t.Fatalf(
						"expected calling `BruteForceMatch()` to find matches of '%s' in '%s' to return %v, got %v",
						tc.pattern,
						tc.text,
						tc.expected,
						got,
					)
				}
			},
		)
	}

}
