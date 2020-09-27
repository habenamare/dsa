package patternmatching

type Match struct {
	Starts int
	Ends   int
}

type Matcher interface {
	Match(pattern string, text string) []Match
}
