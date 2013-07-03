package cmp

// Byte
type Byte byte

// Byte::Cmp
func (a Byte) Cmp(b_ interface{}) int {
	return CompareByte(a, b_)
}

// CompareByte
func CompareByte(a_, b_ interface{}) int {
	a := a_.(Byte)
	b := b_.(Byte)
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
	//panic("Unreachable")
}
