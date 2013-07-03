package cmp

// Int
type Int int

// Int::Cmp
func (a Int) Cmp(b_ interface{}) int {
	return CompareInt(a, b_)
}

// CompareInt
func CompareInt(a_, b_ interface{}) int {
	a := a_.(Int)
	b := b_.(Int)
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
