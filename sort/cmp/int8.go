package cmp

// Int8
type Int8 int8

// Int8::Cmp
func (a Int8) Cmp(b_ interface{}) int {
	return CompareInt8(a, b_)
}

// CompareInt8
func CompareInt8(a_, b_ interface{}) int {
	a := a_.(Int8)
	b := b_.(Int8)
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
