package cmp

// Int16
type Int16 int16

// Int16::Cmp
func (a Int16) Cmp(b_ interface{}) int {
	return CompareInt16(a, b_)
}

// CompareInt16
func CompareInt16(a_, b_ interface{}) int {
	a := a_.(Int16)
	b := b_.(Int16)
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
