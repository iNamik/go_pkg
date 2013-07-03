package cmp

// Int64
type Int64 int64

// Int64::Cmp
func (a Int64) Cmp(b_ interface{}) int {
	return CompareInt64(a, b_)
}

// CompareInt64
func CompareInt64(a_, b_ interface{}) int {
	a := a_.(Int64)
	b := b_.(Int64)
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
