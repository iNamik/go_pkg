go_pkg
=======

**Go Packages**


ABOUT
-----

go_pkg is a convenient place to store Go packages that are too small to warrant their own github project.


PACKAGES
--------

 * **bufio/bleeder/Bleeder**

 Bleeder provides a mechanism for seemlessly exhausting the buffered content of
 a bufio.Reader, then delegating further read requests to a separate reader.

 * **debug/assert**

 Assert provides a simple assert mechanism.

 * **debug/ping**

 Ping provides a simple debug-printing mechanism.

 * **encoding/binary**

 Binary has constants and functions for approximating binary literals.

 * **runes/IndexRune()**

 IndexRune() helps you deterimine if a rune exists in an array of runes.


INSTALL
-------

The packages are built using the Go tool.  Assuming you have correctly set the
$GOPATH variable, you can run the folloing command:

	go get github.com/iNamik/go_pkg


AUTHORS
-------

 * David Farrell
