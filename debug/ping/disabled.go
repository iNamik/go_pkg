// +build disable_ping

package ping

const PING bool = false

func Ping(a ...interface{}) {}

func Pingf(format string, a ...interface{}) {}
