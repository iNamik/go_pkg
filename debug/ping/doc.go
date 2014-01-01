/*
Ping provides a simple debug-printing mechanism.

Below is a simple example, using Ping to print "Hello, world"

*NOTE*  Ping was placed into its own sub-folder to make it safer to import into
        your global scope (i.e. 'import . "")

Example:

package main

import . "github.com/iNamik/go_pkg/debug/ping"

func main() {
	Ping("Hello, world")
}


Ping is enabled by default.  To disable, build your code with the
'disable_ping' tag, i.e.

	go build -tags disable_ping

You can also quickly disable pings by modifying your import like this:

	import . "github.com/iNamik/go_pkg/debug/ping/disabled" // disabled

You can use the PING constant to check if pings are enabled:

	if PING {
		Ping("Hello, world")
	}

When disabled, the Ping functions reduce to:

	func Ping(a ...interface{}) {}

	func Pingf(format string, a ...interface{}) {}

*/
package ping
