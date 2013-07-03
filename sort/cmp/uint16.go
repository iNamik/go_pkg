package cmp

// Uint16
type Uint16 uint16

// Uint16::Cmp
func (a Uint16) Cmp(b_ interface{}) int {
	return CompareUint16(a, b_)
}

// CompareUint16
func CompareUint16(a_, b_ interface{}) int {
	a := a_.(Uint16)
	b := b_.(Uint16)
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
