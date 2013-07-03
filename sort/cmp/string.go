package cmp

// String
type String string

// String::Cmp
func (a String) Cmp(b_ interface{}) int {
	return CompareString(a, b_)
}

// CompareString
// NOTE: I'm sure this is considered a naive implementation
func CompareString(a_, b_ interface{}) int {
	a := a_.(String)
	b := b_.(String)
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
