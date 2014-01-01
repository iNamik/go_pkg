/*

Package maybe implements an option type that may either contain a single value or nothing.

Impure
------

Although inspired by the maybe monad, this implementation is less focused on functional purity,
and more focused on enabling composition with idiomatic go.

New() is used in place of Just(), and contains convenience logic to make working with maybe easier.

The composition methods work with the "comma ok" and "comma err" idioms common in go.

Errors
------

Unlike a pure maybe, in this package a nothing can carry an error.  Granted, it may be difficult
(or impossible) to know where, in a composition, and error has happened, having access to the error
may still be useful.  And since we're supporting the "comma err" idiom, it just makes sense to have it.


Reflection
----------

This package uses reflection in order to deal with functions having varying signatures.

*/
package maybe

import (
	"fmt"
	"reflect"
)

/*********************************************************************
 ** Types & Vars
 *********************************************************************/

// maybe.T
type T struct {
	thing interface{}
	err   error
}

type (
	F_MAP    func(*T) *T
	F_FILTER func(*T) bool
)

// nothing represents a nothing without an error.
// Used internally to minimize memory allocation
var nothing *T = &T{nil, nil}

/*********************************************************************
 ** maybe.T Constructors
 *********************************************************************/

// New creates a maybe with a non-nil value.
// Reduces to Nothing() if value is nil.
// Returns value unmodified if it is already a maybe.
func New(i interface{}) *T {
	// Nothing?
	if i == nil {
		return nothing
	}
	// Already a maybe?
	if m, ok := i.(*T); ok {
		return m
	}
	return &T{i, nil}
}

// Nothing returns a maybe with no value, and no error.
// Internally, the system keeps a single copy of a nothing, and
// uses it whenever a nothing is needed.
func Nothing() *T { return nothing }

// Err creates a nothing containing an error.
// Reduces to Nothing() if err is nil.
func Err(err error) *T {
	if err == nil {
		return nothing
	}
	return &T{nil, err}
}

// ComposeMap accepts a list of map functions, and returns a new function
// that composes the supplied functions in reverse order.
// i.e. ComposeMap(f1, f2) => f(x) maybe { return f1(f2(x)) }
func ComposeMap(funcs ...interface{}) func(*T) *T {
	return func(m *T) *T {
		// Call functions in reverse order
		for i := len(funcs) - 1; i >= 0; i-- {
			m = m.Map(funcs[i])
		}
		return m
	}
}

// ComposeFilter accepts a list of filter functions, and returns a new function
// that composes the supplied functions in reverse order.
// i.e. ComposeFilter(f1, f2) => f(x) bool { return f1(f2(x)) }
func ComposeFilter(funcs ...interface{}) func(*T) bool {
	return func(m *T) bool {
		// Call functions in reverse order
		for i := len(funcs) - 1; i >= 0; i-- {
			m = m.Filter(funcs[i])
		}
		return m.thing != nil
	}
}

// Call calls func f, unwrapping any maybe args, and wrapping the result.
// f is not called if any maybe args are nothing
func Call(f interface{}, _args ...interface{}) *T {
	return call(f, "maybe.Call()", _args...)
}

// call
func call(f interface{}, panicFrom string, _args ...interface{}) *T {
	// If we don't have a func, then panic
	if f == nil {
		panic(fmt.Sprint(panicFrom, " expects func"))
	}
	// Lets look at what we have
	fv := reflect.Indirect(reflect.ValueOf(f))
	// If we don't have a func, then panic
	if fv.Type().Kind() != reflect.Func {
		panic(fmt.Sprint(panicFrom, " expects func"))
	}
	// Build arguments
	args := make([]reflect.Value, len(_args))
	for i, arg := range _args {
		// Is arg a Maybe?
		if m, ok := arg.(*T); ok {
			// If arg is nothing, then return nothing
			if m.thing == nil {
				return nothing
			} else {
				arg = m.thing
			}
		}
		args[i] = reflect.ValueOf(arg)
	}
	// Examine the type
	ft := fv.Type()
	// If only 1 return value, then use it
	if ft.NumOut() == 1 {
		// Try to call the function with our unwrapped thing
		// NOTE: This generates crummy error messages
		ret := fv.Call(args)
		return New(ret[0].Interface())

		// If 2 return values, may be comma-ok or comma-error idiom
	} else if ft.NumOut() == 2 {
		r2t := ft.Out(1)
		switch r2t.Kind() {
		// Bool: Comma-Ok idiom
		case reflect.Bool:
			ret := fv.Call(args)
			if ret[1].Bool() == false {
				return nothing
			}
			return New(ret[0].Interface())

			// Error: Comma-Error idiom
		case reflect.Interface:
			// This feels like a hack, but don't know a better way
			if r2t.PkgPath() == "" && r2t.Name() == "error" {
				ret := fv.Call(args)
				if ret[1].IsNil() == false {
					return Err(ret[1].Interface().(error))
				}
				return New(ret[0].Interface())
			}
		}
	}
	panic(fmt.Sprint(panicFrom, " expects return types to be one of: (any), (any, bool), (any, error)"))
}

/*********************************************************************
 ** maybe.T Functions
 *********************************************************************/

// T::Get returns value,true if the maybe contains something, <undefined>,false otherwise
func (m *T) Get() (interface{}, bool) {
	return m.thing, m.thing != nil
}

// T::GetOrNil returns the maybe's value if it has something, or nill
func (m *T) GetOrNil() interface{} {
	return m.thing
}

// T::GetOrElse returns the maybe's value if it has something, else the passed-in value
func (m *T) GetOrElse(e interface{}) interface{} {
	if m.thing == nil {
		return e
	}
	return m.thing
}

// T::Err returns the maybe's error, or nil if no error
func (m *T) Err() error {
	return m.err
}

// T::Map
func (m *T) Map(f interface{}) *T {
	// If we're already nothing, then we're done
	if m.thing == nil {
		return m
	}
	// If f is composable, use it
	if fc, ok := f.(func(*T) *T); ok {
		return New(fc(m)) // Unknown func may return nil, so need to check with New()
	}
	return call(f, "maybe.T.Map()", m.thing)
}

// T::Filter
func (m *T) Filter(f interface{}) *T {
	// If we're already nothing, then we're done
	if m.thing == nil {
		return m
	}
	// If f is a maybe filter, use it
	if ff, ok := f.(func(*T) bool); ok {
		if ff(m) == false {
			return nothing
		}
		return m
	}
	// If f is nil
	if f == nil {
		panic("maybe.T.Filter() expects func with single return value of type bool")
	}
	// Lets look at what we have
	fv := reflect.Indirect(reflect.ValueOf(f))
	// Filter expects func with single return value of type bool
	if fv.Type().Kind() != reflect.Func || fv.Type().NumOut() != 1 || fv.Type().Out(0).Kind() != reflect.Bool {
		panic("maybe.T.Filter() expects func with single return value of type bool")
	}
	// Try to call the function with our unwrapped thing
	// NOTE: This generates crummy error messages
	ret := fv.Call([]reflect.Value{reflect.ValueOf(m.thing)})
	if ret[0].Bool() == false {
		return nothing
	}
	return m
}
