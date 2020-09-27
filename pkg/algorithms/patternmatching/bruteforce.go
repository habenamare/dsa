package patternmatching

func BruteForceMatch(pattern string, text string) []Match {
	matches := []Match{}

	for i := 0; i < len(text)-len(pattern); i++ {
		starts := i
		ends := i

		for j := 0; j < len(pattern); j++ {
			if pattern[j] != text[i+j] {
				break
			}

			ends = i + j
		}

		if ends-starts == len(pattern)-1 {
			matches = append(matches, Match{
				Starts: starts,
				Ends:   ends,
			})
		}
	}

	return matches
}
