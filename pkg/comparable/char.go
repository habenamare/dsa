package comparable

type Char struct {
	value rune
}

func NewChar(value rune) *Char {
	return &Char{value: value}
}

func (c *Char) Value() interface{} {
	return c.value
}

func (c *Char) Compare(otherChar Comparable) int {
	characterDiff := c.Value().(rune) - otherChar.Value().(rune)

	return int(characterDiff)
}
