// +build !disable_ping

package ping

import (
	"fmt"
	"runtime"
)

const PING bool = true

func Ping(a ...interface{}) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		fmt.Printf("%s:%d: ", file, line)
		if len(a) == 0 {
			fmt.Println("PING")
		} else {
			fmt.Println(a...)
		}
	}
}

func Pingf(format string, a ...interface{}) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		fmt.Printf("%s:%d: ", file, line)
		fmt.Printf(format, a...)
		fmt.Println()
	}
}
