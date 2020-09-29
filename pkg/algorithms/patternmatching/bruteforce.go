package patternmatching

func BruteForceMatch(pattern string, text string) Matches {
	matches := []Match{}

	for i := 0; i < len(text)-(len(pattern)-1); i++ {
		starts, ends := i, i

		matchFound := true
		for j := 0; j < len(pattern); j++ {
			if pattern[j] != text[i+j] {
				matchFound = false
				break
			}

			ends = i + j
		}

		if matchFound {
			matches = append(matches, Match{
				Starts: starts,
				Ends:   ends,
			})
		}
	}

	return matches
}
