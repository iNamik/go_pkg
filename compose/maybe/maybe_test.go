package maybe_test

import (
	"errors"
	//"fmt"
	"github.com/iNamik/go_pkg/compose/maybe"
	//"github.com/iNamik/go_ty"
	"testing"
)

/*********************************************************************
 ** Test Variables / Constants
 *********************************************************************/

const (
	ELSE             = "else"
	ERROR_MSG        = "error message"
	FILTER_PANIC     = "filter_panic"
	MAP_PANIC        = "map_panic"
	MAPPED           = "mapped"
	SOMETHING        = "something"
	FILTER_FUNC_ERR  = "maybe.T.Filter() expects func with single return value of type bool"
	MAP_FUNC_ERR     = "maybe.T.Map() expects func"
	MAP_FUNC_RET_ERR = "maybe.T.Map() expects return types to be one of: (any), (any, bool), (any, error)"
	CALL_FUNC_ERR    = "maybe.Call() expects func"
)

var nothing = maybe.Nothing()

/*********************************************************************
 ** Test Assert Functions
 *********************************************************************/

// assertNothing
func assertNothing(t *testing.T, m *maybe.T) {
	if m != nothing {
		t.Fatal("AssertNothing: m is not equal to maybe.Nothing()")
	}
	_, ok := m.Get()
	if ok {
		t.Fatal("assertNothing: m.Get() returned true")
	}
	err := m.Err()
	if err != nil {
		t.Fatalf("assertNothing: m.Err() returned non-nil")
	}
}

// assertSomething
func assertSomething(t *testing.T, m *maybe.T) {
	if m == nothing {
		t.Fatal("AssertSomething: m is equal to maybe.Nothing()")
	}
	g, ok := m.Get()
	if !ok {
		t.Fatal("assertSomething: m.Get() returned false")
	}
	if g != SOMETHING {
		t.Fatalf("assertSomething: m.Get() returned '%v' instead of '%s'", g, SOMETHING)
	}
	err := m.Err()
	if err != nil {
		t.Fatalf("assertSomething: m.Err() returned non-nil")
	}
}

// assertSomethingValue
func assertSomethingValue(t *testing.T, m *maybe.T, v interface{}) {
	if m == nothing {
		t.Fatal("AssertSomethingValue: m is equal to maybe.Nothing()")
	}
	g, ok := m.Get()
	if !ok {
		t.Fatal("assertSomethingValue(): m.Get() returned false")
	}
	if g != v {
		t.Fatalf("assertSomethingValue(): m.Get() returned '%v' instead of '%v'", g, v)
	}
	err := m.Err()
	if err != nil {
		t.Fatalf("assertSomethingValue: m.Err() returned non-nil")
	}
}

// assertGetOrNil
func assertGetOrNil(t *testing.T, m *maybe.T, v interface{}) {
	g := m.GetOrNil()
	if g != v {
		t.Fatalf("assertGetOrNil: m.Get() returned '%v' instead of '%v'", g, v)
	}
}

// assertGetOrElse
func assertGetOrElse(t *testing.T, m *maybe.T, e string, v string) {
	g := m.GetOrElse(e)
	if g != v {
		t.Fatalf("assertGitOrElse: m.Get() returned '%v' instead of '%s'", g, v)
	}
}

// assertError
func assertError(t *testing.T, m *maybe.T, msg string) {
	if m == nothing {
		t.Fatal("AssertError: m is equal to maybe.Nothing()")
	}
	_, ok := m.Get()
	if ok {
		t.Fatal("assertError: m.Get() returned true")
	}
	err := m.Err()
	if err == nil {
		t.Fatalf("assertError: m.Err() returned nil")
	}
	if err.Error() != msg {
		t.Fatalf("assertError: m.Err() returned '%s' instead of '%s'", err.Error(), msg)
	}
}

// assertPanic
func assertPanic(t *testing.T, f func(), msg string) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("assertPanic: did not generate panic()")
		} else if r != msg {
			t.Fatalf("assertPanic: recover() recieved message '%s' instead of '%s'", r, msg)
		}
	}()
	f()
}

// assertTrue
func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Fatal("assertTrue: b is false")
	}
}

// assertFalse
func assertFalse(t *testing.T, b bool) {
	if b {
		t.Fatal("assertTrue: b is true")
	}
}

/*********************************************************************
 ** Helper Functions
 *********************************************************************/

// MapEcho
func MapEcho(m *maybe.T) *maybe.T {
	return m
}

// MapSomething
func MapSomething(m *maybe.T) *maybe.T {
	return maybe.New(MAPPED)
}

// MapNothing
func MapNothing(m *maybe.T) *maybe.T {
	return maybe.Nothing()
}

// MapPanic
func MapPanic(m *maybe.T) *maybe.T {
	panic(MAP_PANIC)
}

// MapAny_Panic
func MapAny_Panic(i interface{}) interface{} {
	panic(MAP_PANIC)
}

// Map_NoReturn
func Map_NoReturn(i interface{}) {
	panic(MAP_PANIC)
}

// MapAny_Echo
func MapAny_Echo(i interface{}) interface{} {
	return i
}

// MapAny_Nil
func MapAny_Nil(i interface{}) interface{} {
	return nil
}

// MapOk_Panic
func MapOk_Panic(i interface{}) (interface{}, bool) {
	panic(MAP_PANIC)
}

// MapOk_True
func MapOk_True(i interface{}) (interface{}, bool) {
	return i, true
}

// MapOk_False
func MapOk_False(i interface{}) (interface{}, bool) {
	return i, false
}

// MapErr_Panic
func MapErr_Panic(i interface{}) (interface{}, error) {
	panic(MAP_PANIC)
}

// MapErr_Nil
func MapErr_Nil(i interface{}) (interface{}, error) {
	return i, nil
}

// MapErr_Err
func MapErr_Err(i interface{}) (interface{}, error) {
	return i, errors.New(ERROR_MSG)
}

// FilterTrue
func FilterTrue(m *maybe.T) bool {
	return true
}

// FilterFalse
func FilterFalse(m *maybe.T) bool {
	return false
}

// FilterPanic
func FilterPanic(m *maybe.T) bool {
	panic(FILTER_PANIC)
}

// FilterAny_Panic
func FilterAny_Panic(i interface{}) bool {
	panic(FILTER_PANIC)
}

// FilterAny_True
func FilterAny_True(i interface{}) bool {
	return true
}

// FilterAny_False
func FilterAny_False(i interface{}) bool {
	return false
}

// Filter_NoReturn
func Filter_NoReturn(i interface{}) {
	panic(FILTER_PANIC)
}

// NoArgs_Something
func NoArgs_Something() interface{} {
	return SOMETHING
}

// NoArgs_Nil
func NoArgs_Nil() interface{} {
	return nil
}

// MultiArgs_Something
func MultiArgs_Something(args ...interface{}) interface{} {
	return SOMETHING
}

// MultiArgs_Last
func MultiArgs_Last(args ...interface{}) interface{} {
	return args[len(args)-1]
}

// MultiArgs_Nil
func MultiArgs_Nil(args ...interface{}) interface{} {
	return nil
}

// increment
func increment(i int) int {
	return i + 1
}

// double
func double(i int) int {
	return i + i
}

// square
func square(i int) int {
	return i * i
}

// add
func add(a, b int) int {
	return a + b
}

/*********************************************************************
 ** Test Functions: maybe.Nothing()
 *********************************************************************/

// TestNothing
func TestNothing(t *testing.T) {
	m := maybe.Nothing()
	assertNothing(t, m)
}

/*********************************************************************
 ** Test Functions: maybe.Err()
 *********************************************************************/

// TestError_Nil
func TestError_Nil(t *testing.T) {
	m := maybe.Err(nil)
	assertNothing(t, m)
}

// TestError_Error
func TestError_Error(t *testing.T) {
	m := maybe.Err(errors.New(ERROR_MSG))
	assertError(t, m, ERROR_MSG)
}

/*********************************************************************
 ** Test Functions: maybe.New()
 *********************************************************************/

// TestNew_Nil
func TestNew_Nil(t *testing.T) {
	m := maybe.New(nil)
	assertNothing(t, m)
}

// TestNew_Something
func TestNew_Something(t *testing.T) {
	m := maybe.New(SOMETHING)
	assertSomething(t, m)
}

// TestNew_Maybe_Nothing
func TestNew_Maybe_Nothing(t *testing.T) {
	m := maybe.New(nothing)
	assertNothing(t, m)
}

// TestNew_Maybe_Something
func TestNew_Maybe_Something(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := maybe.New(_m)
	assertSomething(t, m)
}

/*********************************************************************
 ** Test Functions: maybe.T.GetOrNil() & maybe.T.GetOrElse()
 *********************************************************************/

// TestGetOrNil_Nothing
func TestGetOrNil_Nothing(t *testing.T) {
	m := maybe.Nothing()
	assertGetOrNil(t, m, nil)
}

// TestGetOrNil_Something
func TestGetOrNil_Something(t *testing.T) {
	m := maybe.New(SOMETHING)
	assertGetOrNil(t, m, SOMETHING)
}

// TestGetOrElse_Nothing
func TestGetOrElse_Nothing(t *testing.T) {
	m := maybe.Nothing()
	assertGetOrElse(t, m, ELSE, ELSE)
}

// TestGetOrElse_Something
func TestGetOrElse_Something(t *testing.T) {
	m := maybe.New(SOMETHING)
	assertGetOrElse(t, m, ELSE, SOMETHING)
}

/*********************************************************************
 ** Test Functions: maybe.T.Map()
 *********************************************************************/

// TestMap_Nothing_Nil
func TestMap_Nothing_Nil(t *testing.T) {
	_m := maybe.Nothing()
	m := _m.Map(nil)
	assertNothing(t, m)
}

// TestMap_Nothing_Panic
func TestMap_Nothing_Panic(t *testing.T) {
	_m := maybe.Nothing()
	m := _m.Map(MapPanic)
	assertNothing(t, m)
}

// TestMap_Something_Nil
func TestMap_Something_Nil(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Map(nil) }
	assertPanic(t, f, MAP_FUNC_ERR)
}

// TestMap_Something_Panic
func TestMap_Something_Panic(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Map(MapPanic) }
	assertPanic(t, f, MAP_PANIC)
}

// TestMap_Something_NoReturn
func TestMap_Something_NoReturn(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Map(Map_NoReturn) }
	assertPanic(t, f, MAP_FUNC_RET_ERR)
}

// TestMap_Something_NonFunc
func TestMap_Something_NonFunc(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Map(0) } // int 0 is not a func
	assertPanic(t, f, MAP_FUNC_ERR)
}

// TestMap_Something_Nothing
func TestMap_Something_Nothing(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapNothing)
	assertNothing(t, m)
}

// TestMap_Something_Something
func TestMap_Something_Something(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapSomething)
	assertSomethingValue(t, m, MAPPED)
}

// TestMap_Something_Any_Panic
func TestMap_Something_Any_Panic(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Map(MapAny_Panic) }
	assertPanic(t, f, MAP_PANIC)
}

// TestMap_Something_Any_Nil
func TestMap_Something_Any_Nil(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapAny_Nil)
	assertNothing(t, m)
}

// TestMap_Something_Any_Something
func TestMap_Something_Any_Something(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapAny_Echo)
	assertSomething(t, m)
}

// TestMap_Something_Ok_Panic
func TestMap_Something_Ok_Panic(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Map(MapOk_Panic) }
	assertPanic(t, f, MAP_PANIC)
}

// TestMap_Something_Ok_False
func TestMap_Something_Ok_False(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapOk_False)
	assertNothing(t, m)
}

// TestMap_Something_OK_True
func TestMap_Something_Ok_True(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapOk_True)
	assertSomething(t, m)
}

// TestMap_Something_Err_Panic
func TestMap_Something_Err_Panic(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Map(MapErr_Panic) }
	assertPanic(t, f, MAP_PANIC)
}

// TestMap_Something_Err_Err
func TestMap_Something_Err_Err(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapErr_Err)
	assertError(t, m, ERROR_MSG)
}

// TestMap_Something_Err_Nil
func TestMap_Something_Err_Nil(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapErr_Nil)
	assertSomething(t, m)
}

// TestMap_Multi_Something_Ok_True_False
func TestMap_Multi_Something_Ok_True_False(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapOk_True).Map(MapOk_False)
	assertNothing(t, m)
}

// TestMap_Multi_Something_Ok_True_True
func TestMap_Multi_Something_Ok_True_True(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Map(MapOk_True).Map(MapOk_True)
	assertSomething(t, m)
}

/*********************************************************************
 ** Test Functions: maybe.T.Filter()
 *********************************************************************/

// TestFilter_Nothing_Nil
func TestFilter_Nothing_Nil(t *testing.T) {
	_m := maybe.Nothing()
	m := _m.Filter(nil)
	assertNothing(t, m)
}

// TestFilter_Nothing_Panic
func TestFilter_Nothing_Panic(t *testing.T) {
	_m := maybe.Nothing()
	m := _m.Filter(FilterPanic)
	assertNothing(t, m)
}

// TestFilter_Something_Nil
func TestFilter_Something_Nil(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Filter(nil) }
	assertPanic(t, f, FILTER_FUNC_ERR)
}

// TestFilter_Something_Panic
func TestFilter_Something_Panic(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Filter(FilterPanic) }
	assertPanic(t, f, FILTER_PANIC)
}

// TestFilter_Something_False
func TestFilter_Something_False(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Filter(FilterFalse)
	assertNothing(t, m)
}

// TestFilter_Something_True
func TestFilter_Something_True(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Filter(FilterTrue)
	assertSomething(t, m)
}

// TestFilter_Something_Any_Panic
func TestFilter_Something_Any_Panic(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Filter(FilterAny_Panic) }
	assertPanic(t, f, FILTER_PANIC)
}

// TestFilter_Something_Any_False
func TestFilter_Something_Any_False(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Filter(FilterAny_False)
	assertNothing(t, m)
}

// TestFilter_Something_Any_True
func TestFilter_Something_Any_True(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Filter(FilterAny_True)
	assertSomething(t, m)
}

// TestFilter_Something_NoReturn
func TestFilter_Something_NoReturn(t *testing.T) {
	m := maybe.New(SOMETHING)
	f := func() { m.Filter(Filter_NoReturn) }
	assertPanic(t, f, FILTER_FUNC_ERR)
}

// TestFilter_Multi_Something_True_False
func TestFilter_Multi_Something_True_False(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Filter(FilterTrue).Filter(FilterFalse)
	assertNothing(t, m)
}

// TestFilter_Multi_Something_True_True
func TestFilter_Multi_Something_True_True(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := _m.Filter(FilterTrue).Filter(FilterTrue)
	assertSomething(t, m)
}

/*********************************************************************
 ** Test Functions: maybe.Call()
 *********************************************************************/

// TestCall_Nil
func TestCall_Nil(t *testing.T) {
	f := func() { maybe.Call(nil, 0) }
	assertPanic(t, f, CALL_FUNC_ERR)
}

// TestCall_Panic
func TestCall_Panic(t *testing.T) {
	f := func() { maybe.Call(MapAny_Panic, 0) }
	assertPanic(t, f, MAP_PANIC)
}

// TestCall_NoArgs
func TestCall_NoArgs(t *testing.T) {
	m := maybe.Call(NoArgs_Something)
	assertSomething(t, m)
}

// TestCall_Something_Echo
func TestCall_Something_Echo(t *testing.T) {
	_m := maybe.New(SOMETHING)
	m := maybe.Call(MapAny_Echo, _m)
	assertSomething(t, m)
}

// TestCall_Something_Double
func TestCall_Something_Double(t *testing.T) {
	_m := maybe.New(3)
	m := maybe.Call(double, _m)
	assertSomethingValue(t, m, 6)
}

// TestCall_Nothing
func TestCall_Nothing(t *testing.T) {
	_m := maybe.Nothing()
	m := maybe.Call(MapAny_Echo, _m)
	assertNothing(t, m)
}

// TestCall_Any
func TestCall_Any(t *testing.T) {
	m := maybe.Call(MapAny_Echo, SOMETHING)
	assertSomething(t, m)
}

// TestCall_Multi_Something_Last
func TestCall_Multi_Something_Last(t *testing.T) {
	arg1 := maybe.New("one")
	arg2 := maybe.New("two")
	m := maybe.Call(MultiArgs_Last, arg1, arg2)
	assertSomethingValue(t, m, "two")
}

// TestCall_Multi_Something_Add
func TestCall_Multi_Something_Add(t *testing.T) {
	arg1 := maybe.New(1)
	arg2 := maybe.New(2)
	m := maybe.Call(add, arg1, arg2)
	assertSomethingValue(t, m, 3)
}

// TestCall_Multi_Nil
func TestCall_Multi_Nil(t *testing.T) {
	arg1 := maybe.New("one")
	arg2 := maybe.Nothing()
	m := maybe.Call(MultiArgs_Last, arg1, arg2)
	assertNothing(t, m)
}

// TestCall_Multi_Mixed_Add
func TestCall_Multi_Mixed_Add(t *testing.T) {
	arg1 := maybe.New(1)
	m := maybe.Call(add, arg1, 2)
	assertSomethingValue(t, m, 3)
}

/*********************************************************************
 ** Test Functions: maybe.ComposeMap()
 *********************************************************************/

// TestComposeMap_Empty_Nothing
func TestComposeMap_Empty_Nothing(t *testing.T) {
	f := maybe.ComposeMap()
	_m := maybe.Nothing()
	m := f(_m)
	assertNothing(t, m)
}

// TestComposeMap_Empty_Something
func TestComposeMap_Empty_Something(t *testing.T) {
	f := maybe.ComposeMap()
	_m := maybe.New(SOMETHING)
	m := f(_m)
	assertSomething(t, m)
}

// TestComposeMap_Increment_Nothing
func TestComposeMap_Increment_Nothing(t *testing.T) {
	f := maybe.ComposeMap(increment)
	_m := maybe.Nothing()
	m := f(_m)
	assertNothing(t, m)
}

// TestComposeMap_Increment_Something
func TestComposeMap_Increment_Something(t *testing.T) {
	f := maybe.ComposeMap(increment)
	_m := maybe.New(1)
	m := f(_m)
	assertSomethingValue(t, m, 2)
}

// TestComposeMap_Increment_Double_Something
func TestComposeMap_Increment_Double_Something(t *testing.T) {
	f := maybe.ComposeMap(double, increment)
	_m := maybe.New(2)
	m := f(_m)
	assertSomethingValue(t, m, 6)
}

// TestComposeMap_Increment_Echo_Square_Something
func TestComposeMap_Increment_Echo_Square_Something(t *testing.T) {
	f := maybe.ComposeMap(square, MapEcho, increment)
	_m := maybe.New(3)
	m := f(_m)
	assertSomethingValue(t, m, 16)
}

/*********************************************************************
 ** Test Functions: maybe.ComposeFilter()
 *********************************************************************/

// TestComposeFilter_Empty_Nothing
func TestComposeFilter_Empty_Nothing(t *testing.T) {
	f := maybe.ComposeFilter()
	m := maybe.Nothing()
	b := f(m)
	assertFalse(t, b)
}

// TestComposeFilter_Empty_Something
func TestComposeFilter_Empty_Something(t *testing.T) {
	f := maybe.ComposeFilter()
	m := maybe.New(SOMETHING)
	b := f(m)
	assertTrue(t, b)
}

// TestComposeFilter_True_Nothing
func TestComposeFilter_True_Nothing(t *testing.T) {
	f := maybe.ComposeFilter(FilterTrue)
	m := maybe.Nothing()
	b := f(m)
	assertFalse(t, b)
}

// TestComposeFilter_True_Something
func TestComposeFilter_True_Something(t *testing.T) {
	f := maybe.ComposeFilter(FilterTrue)
	m := maybe.New(SOMETHING)
	b := f(m)
	assertTrue(t, b)
}

// TestComposeFilter_FalseAny_Something
func TestComposeFilter_FalseAny_Something(t *testing.T) {
	f := maybe.ComposeFilter(FilterAny_False)
	m := maybe.New(SOMETHING)
	b := f(m)
	assertFalse(t, b)
}

// TestComposeFilter_TrueAny_Something
func TestComposeFilter_TrueAny_Something(t *testing.T) {
	f := maybe.ComposeFilter(FilterAny_True)
	m := maybe.New(SOMETHING)
	b := f(m)
	assertTrue(t, b)
}

// TestComposeFilter_True_TrueAny_Something
func TestComposeFilter_True_TrueAny_Something(t *testing.T) {
	f := maybe.ComposeFilter(FilterTrue, FilterAny_True)
	m := maybe.New(SOMETHING)
	b := f(m)
	assertTrue(t, b)
}

// TestComposeFilter_True_FalseAny_Something
func TestComposeFilter_True_FalseAny_Something(t *testing.T) {
	f := maybe.ComposeFilter(FilterTrue, FilterAny_False)
	m := maybe.New(SOMETHING)
	b := f(m)
	assertFalse(t, b)
}
