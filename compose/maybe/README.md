maybe
=====

About
-----

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

