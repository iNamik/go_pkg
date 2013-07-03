package cmp

// Uint
type Uint uint

// Uint::Cmp
func (a Uint) Cmp(b_ interface{}) int {
	return CompareUint(a, b_)
}

// CompareUint
func CompareUint(a_, b_ interface{}) int {
	a := a_.(Uint)
	b := b_.(Uint)
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
