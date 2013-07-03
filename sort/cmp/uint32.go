package cmp

// Uint32
type Uint32 uint32

// Uint32::Cmp
func (a Uint32) Cmp(b_ interface{}) int {
	return CompareUint32(a, b_)
}

// CompareUint32
func CompareUint32(a_, b_ interface{}) int {
	a := a_.(Uint32)
	b := b_.(Uint32)
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
