/*
The testutil package contains a few helper functions and wrappers that work
with Golang's test package. Most of the functions have names similar in form to
for example AssertEqInt(). This example takes an expected integer argument and
a received integer argument and ensures they are equal. This family of
functions effectivelymoves the boiler plate code of if statements away from
test functions and into this library. The test function consumers need only
call one of these functions in fire-and-forget mode. The functions include in
their output to the testing.T object, a simplified stack trace so that it is
easy to find the consuming function that has stimulated the failed test. The
functions take an argument called 'thing'. This is printed as part of the test
failure diagnostics message.
*/
package testutil

import (
	"gopkg.in/yaml.v2"
	"runtime"
	"strings"
	"testing"
)

func AssertEqInt(t *testing.T, actual int, expected int, thing string) {
	if actual == expected {
		return
	}
	t.Errorf("%s is wrong: %v, expected: %v", thing, actual, expected)
	t.Error(briefStackTrace())
}

func AssertEqString(t *testing.T, actual string, expected string, thing string) {
	if actual == expected {
		return
	}
	t.Errorf("%s is wrong: %v, expected: %v", thing, actual, expected)
	t.Error(briefStackTrace())
}

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

func AssertStrContains(t *testing.T, main string, sub string, thing string) {
	if strings.Contains(main, sub) {
		return
	}
	t.Errorf("This string: <%s> does not contain <%s>", main, sub)
	t.Error(briefStackTrace())
}

// Ensure that the given error object is nil.
func AssertNilErr(t *testing.T, err error, thing string) {
	if err == nil {
		return
	}
	t.Errorf("Error generated: %v", err.Error())
	t.Error(briefStackTrace())
}

// Ensure that the string returned by the given error's Error() method,
// contains the given sub string. In other words - did the right error message
// get produced.
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

/*
The function briefStackTrace() generates a stack trace and then reduces the
data therein to only the lines that include source code line numbers.
*/
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
