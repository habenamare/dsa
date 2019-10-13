package comparable

type Integer struct {
	value int
}

func NewInteger(value int) *Integer {
	return &Integer{value: value}
}

func (i *Integer) Value() interface{} {
	return i.value
}

func (i *Integer) Compare(otherInteger Comparable) int {
	return i.Value().(int) - otherInteger.Value().(int)
}
