package cmp

// Float32
type Float32 float32

// Float32::Cmp
func (a Float32) Cmp(b_ interface{}) int {
	return CompareFloat32(a, b_)
}

// CompareFloat32
func CompareFloat32(a_, b_ interface{}) int {
	a := a_.(Float32)
	b := b_.(Float32)
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
