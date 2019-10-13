package comparable

type Comparable interface {
	Compare(Comparable) int
	Value() interface{}
}
