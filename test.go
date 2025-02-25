package termstatus

import (
	"fmt"
	"reflect"
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	tb.Helper()
	if !condition {
		tb.Fatalf("\033[31m"+msg+"\033[39m\n\n", v...)
	}
}

// OK fails the test if an err is not nil.
func OK(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("\033[31munexpected error: %+v\033[39m\n\n", err)
	}
}

// Equals fails the test if exp is not equal to act.
// msg is optional message to be printed, first param being format string and rest being arguments.
func Equals(tb testing.TB, exp, act interface{}, msgs ...string) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		var msgString string
		length := len(msgs)
		if length == 1 {
			msgString = msgs[0]
		} else if length > 1 {
			args := make([]interface{}, length-1)
			for i, msg := range msgs[1:] {
				args[i] = msg
			}
			msgString = fmt.Sprintf(msgs[0], args...)
		}
		tb.Fatalf("\033[31m\n\n\t"+msgString+"\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", exp, act)
	}
}
