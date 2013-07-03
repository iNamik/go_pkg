package cmp

// Uint64
type Uint64 uint64

// Uint64::Cmp
func (a Uint64) Cmp(b_ interface{}) int {
	return CompareUint64(a, b_)
}

// CompareUint64
func CompareUint64(a_, b_ interface{}) int {
	a := a_.(Uint64)
	b := b_.(Uint64)
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
