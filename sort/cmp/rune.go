package cmp

// Rune
type Rune rune

// Rune::Cmp
func (a Rune) Cmp(b_ interface{}) int {
	return CompareRune(a, b_)
}

// CompareRune
func CompareRune(a_, b_ interface{}) int {
	a := a_.(Rune)
	b := b_.(Rune)
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
