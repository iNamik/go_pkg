/*
Package assert adds conditionally-compiled* asserts to your project.

The package is designed to be imported into your main scope for convenient access:

	import . "github.com/iNamik/go_pkg/debug/assert"

	Assert(true)
	Assert(false) // generates panic

Assert is enabled by default.  To disable them, build your code with the
'disable_assert' tag, i.e.

	go build -tags disable_assert

You can also quickly disable your asserts by modifying your import like this:

	import . "github.com/iNamik/go_pkg/debug/assert/disabled" // disabled

You can use the ASSERT constant to check if asserts are enabled.  This
may allow you to avoid 'heavy' asserts when they are disabled:

	if ASSERT {
		Assert( / * heavy conditional * / )
	}

Or perhaps:

	Assert( ASSERT && / * heavy conditional * /)


* Conditional Compilation *

When disabled, the Assert() function reduces to :

	func Assert(b bool) {}

An assumption could be made that the compiler will detect this empty function,
that has no action and no side-effects, and remove the asserts.

As of Go 1.1, this does not appear to be the case, but I have hopes that future
versions will detect this condition.


* Side Effects *

Your best bet at future versions detecting/removing disabled Asserts is to ensure
that your asserts have no side effects.

This means ensuring that your asserts do not change state.

It should go without saying that you should avoid assignments in asserts for
many reasons, state being just one.

Additionally, the easiest way to avoid side effects is to avoid having any
function calls within the assert.  The compiler may not (or worse may never)
be able to prove that your function call has no side effects, which may result
in your assert never being optimized out.

*/
package assert
