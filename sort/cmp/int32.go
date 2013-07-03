package cmp

// Int32
type Int32 int32

// Int32::Cmp
func (a Int32) Cmp(b_ interface{}) int {
	return CompareInt32(a, b_)
}

// CompareInt32
func CompareInt32(a_, b_ interface{}) int {
	a := a_.(Int32)
	b := b_.(Int32)
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
