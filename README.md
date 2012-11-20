go_pkg
=======

**Go Packages**


ABOUT
-----

go_pkg is a convenient place to store Go packages that are too small to warrant their own github project.


PACKAGES
--------

 * **bufio/Bleeder**

 Bleeder provides a mechanism for seemlessly exhausting the buffered content of
 a bufio.Reader, then delegating further read requests to a separate reader.

 * **debug/Ping**

 Ping provides a simple debug-printing mechanism


INSTALL
-------

The packages are built using the Go tool.  Assuming you have correctly set the
$GOPATH variable, you can run the folloing command:

	go get github.com/iNamik/go_pkg


AUTHORS
-------

 * David Farrell
