package util

import (
	"runtime"
	"strings"
	"testing"
)

/*
The function AssertEqInt32 is a helper function for the golang test package.  It
is syntax sugar for asserting that a given int32 matches an expected value.  When
the assertion does not hold, it sends a message to the injected testing.T object
which includes a simplified stack trace with line numbers. The function parameter
'thing' gets used in the message generated as a noun for the thing that is wrong.
*/
func AssertEqInt32(t *testing.T, actual int32, expected int32, thing string) {
	if actual == expected {
		return
	}
	t.Errorf("%s is wrong: %v, expected: %v", thing, actual, expected)
	t.Error(briefStackTrace())
}

/*
The function AssertStrContains is a helper function for the golang test package.
It is syntax sugar for asserting that a given string contains a given sub string.
When the assertion does not hold, it sends a message to the injected testing.T
object which includes a simplified stack trace with line numbers. The function
parameter 'thing' gets used in the message generated as a noun for the thing that
is wrong.
*/
func AssertStrContains(t *testing.T, main string, sub string, thing string) {
	if strings.Contains(main, sub) {
		return
	}
	t.Errorf("This string: <%s> does not contain <%s>", main, sub)
	t.Error(briefStackTrace())
}

/*
The function AssertNilErr is a helper function for the golang test package.  It
is syntax sugar for asserting that a given error value is not nil.  When the
assertion does not hold, it sends a message to the injected testing.T object
which includes a simplified stack trace with line numbers. The function parameter
'thing' gets used in the message generated as a noun for the thing that is wrong.
*/
func AssertNilErr(t *testing.T, err error, thing string) {
	if err == nil {
		return
	}
	t.Errorf("Error generated: %v", err.Error())
	t.Error(briefStackTrace())
}

// The function briefStackTrace() generates a stack trace and then reduces the
// data therein to only the lines that include source code line numbers.
func briefStackTrace() (briefTrace string) {
	var buf = make([]byte, 10000) // has to be big enough
	written := runtime.Stack(buf, false)
	fullTrace := string(buf[0:written])
	lines := strings.Split(fullTrace, "\n")
	keep := []string{}
	for _, line := range lines {
		if strings.Contains(line, ".go:") {
			keep = append(keep, line)
		}
	}
	return strings.Join(keep, "\n")
}
