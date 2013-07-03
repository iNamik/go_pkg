package cmp

// Result - Optional, can be used to make function/type signatures more clear
// Example:
//	func my_cmp_func(a, b interface{}) cmp.Result
type Result int

// Compare Results
const (
	LT = -1 // Less-Than
	EQ = 0  // Equal
	GT = 1  // Greater-Than
)

// T allows objects to export a comparator function
type T interface {
	Cmp(interface{}) int
}

// F allows you to pass functions that perform comparisons
type F func(a, b interface{}) int

// TF allows you to cast a func/closure into cmp.T interface.
// Example:
//	var key cmp.T = cmp.TF(func(v interface{}) int { /*...*/ })
type TF func(interface{}) int

// TF::Cmp fulfills the cmp.T interface.
// NOTE: Yes you can treat functions as interaces, but I can't find documented source of the feature.
func (f TF) Cmp(v interface{}) int {
	return f(v)
}
