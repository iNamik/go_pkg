// +build !disable_assert

package assert

import (
	"fmt"
	"runtime"
)

const ASSERT bool = true

func Assert(b bool) {
	if !b {
		if _, file, line, ok := runtime.Caller(1); ok {
			panic(fmt.Sprintf("%s:%d: ASSERTION FAILED", file, line))
		} else {
			panic("ASSERTION FAILED")
		}
	}
}
