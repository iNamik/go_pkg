/*
Ping provides a simple debug-printing mechanism.

Below is a simple example, using Pingf to print "Hello, world"

*NOTE*  Ping was placed into its own sub-folder to make it safer to import into
        your global scope (i.e. 'import . "")

Example:

package main

import . "github.com/iNamik/go_pkg/debug/ping"

func main() {
	Pingf("Hello, world")
}

*/
package ping

import (
	"fmt"
	"runtime"
)

func Ping() {
	_, _, line, ok := runtime.Caller(1)

	if ok {
		fmt.Printf("%d: PING\n", line)
	}
}

func Pingf(format string, a ...interface{}) {
	_, _, line, ok := runtime.Caller(1)

	if ok {
		fmt.Printf("%d: ", line)
		fmt.Printf(format, a...)
		fmt.Println()
	}
}
