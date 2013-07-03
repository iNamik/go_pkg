package cmp

// Uint8
type Uint8 uint8

// Uint8::Cmp
func (a Uint8) Cmp(b_ interface{}) int {
	return CompareUint8(a, b_)
}

// CompareUint8
func CompareUint8(a_, b_ interface{}) int {
	a := a_.(Uint8)
	b := b_.(Uint8)
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
