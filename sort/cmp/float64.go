package cmp

// Float64
type Float64 float64

// Float64::Cmp
func (a Float64) Cmp(b_ interface{}) int {
	return CompareFloat64(a, b_)
}

// CompareFloat64
func CompareFloat64(a_, b_ interface{}) int {
	a := a_.(Float64)
	b := b_.(Float64)
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
