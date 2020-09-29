package patternmatching

type Match struct {
	Starts int
	Ends   int
}

func NewMatch(starts int, ends int) Match {
	return Match{
		Starts: starts,
		Ends:   ends,
	}
}

type Matches []Match

func (matches Matches) Equals(anotherM Matches) bool {
	if len(matches) != len(anotherM) {
		return false
	}

	for i, m := range matches {
		if m.Starts != anotherM[i].Starts || m.Ends != anotherM[i].Ends {
			return false
		}
	}

	return true
}

type Matcher interface {
	Match(pattern string, text string) []Match
}
