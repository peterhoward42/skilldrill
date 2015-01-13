package testutil

import (
	"gopkg.in/yaml.v2"
	"runtime"
	"strings"
	"testing"
)

/*
The function AssertEqInt is a helper function for the golang test package.  It
is syntax sugar for asserting that a given int matches an expected value.  When
the assertion does not hold, it sends a message to the injected testing.T object
which includes a simplified stack trace with line numbers. The function parameter
'thing' gets used in the message generated as a noun for the thing that is wrong.
*/
func AssertEqInt(t *testing.T, actual int, expected int, thing string) {
	if actual == expected {
		return
	}
	t.Errorf("%s is wrong: %v, expected: %v", thing, actual, expected)
	t.Error(briefStackTrace())
}

/*
The function AssertEqString is a helper function for the golang test package.  It
is syntax sugar for asserting that a given string matches an expected value.  When
the assertion does not hold, it sends a message to the injected testing.T object
which includes a simplified stack trace with line numbers. The function parameter
'thing' gets used in the message generated as a noun for the thing that is wrong.
*/
func AssertEqString(t *testing.T, actual string, expected string, thing string) {
	if actual == expected {
		return
	}
	t.Errorf("%s is wrong: %v, expected: %v", thing, actual, expected)
	t.Error(briefStackTrace())
}

/*
The function AssertEqSliceInt() is a helper function for the golang test package.
It is syntax sugar for asserting that a given slice of integers matches an
expected value.  When the assertion does not hold, it sends a message to the
injected testing.T object which includes a simplified stack trace with line
numbers. The function parameter 'thing' gets used in the message generated as a
noun for the thing that is wrong.
*/
func AssertEqSliceInt(t *testing.T, actual []int, expected []int, thing string) {
	act, _ := yaml.Marshal(actual)
	actualStr := string(act)
	exp, _ := yaml.Marshal(expected)
	expectedStr := string(exp)

	if actualStr == expectedStr {
		return
	}
	t.Errorf("%s is wrong: %v, expected: %v", thing, actualStr, expectedStr)
	t.Error(briefStackTrace())
}

/*
The function AssertEqSliceString() is a helper function for the golang test
package.  It is syntax sugar for asserting that a given slice of strings matches
an expected value.  When the assertion does not hold, it sends a message to the
injected testing.T object which includes a simplified stack trace with line
numbers. The function parameter 'thing' gets used in the message generated as a
noun for the thing that is wrong.
*/
func AssertEqSliceString(t *testing.T, actual []string, expected []string,
	thing string) {
	act, _ := yaml.Marshal(actual)
	actualStr := string(act)
	exp, _ := yaml.Marshal(expected)
	expectedStr := string(exp)
	if actualStr == expectedStr {
		return
	}
	t.Errorf("%s is wrong: %v, expected: %v", thing, actualStr, expectedStr)
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
is syntax sugar for asserting that a given error value is nil.  When the
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

/*
The function AssertErrGenerated is a helper function for the golang test package.
It is syntax sugar for asserting that a given error value is non-nil and that the
error message includes the given substring.  When the assertion does not hold, it
sends a message to the injected testing.T object which includes a simplified
stack trace with line numbers. The function parameter 'thing' gets used in the
message generated as a noun for the thing that is wrong.
*/
func AssertErrGenerated(t *testing.T, err error, substring string,
	thing string) {
	if err == nil {
		t.Errorf("%s: Error was not generated", thing)
		t.Error(briefStackTrace())
	}
	if !strings.Contains(err.Error(), substring) {
		t.Errorf("%s: Wrong error content: %s, expected to contain <%s>", thing,
			err.Error(), substring)
		t.Error(briefStackTrace())
	}
}

/*
The function AssertTrue is a helper function for the golang test package.  It
is syntax sugar for asserting that a given bool value is true.  When the
assertion does not hold, it sends a message to the injected testing.T object
which includes a simplified stack trace with line numbers. The function parameter
'thing' gets used in the message generated as a noun for the thing that is wrong.
*/
func AssertTrue(t *testing.T, shouldBeTrue bool, thing string) {
	if shouldBeTrue == true {
		return
	}
	t.Errorf("Boolean wrong: %v", thing)
	t.Error(briefStackTrace())
}

func AssertFalse(t *testing.T, shouldBeFalse bool, thing string) {
    shouldBeTrue := !shouldBeFalse
    AssertTrue(t, shouldBeTrue, thing)
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
