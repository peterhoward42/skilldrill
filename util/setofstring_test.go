package util

import (
	"testing"
)

func TestSetOfString(t *testing.T) {
	set := NewSetOfString()

	// Do some adding, including duplicate additions
	set.Add("A")
	set.Add("B")
	set.Add("C")
	set.Add("C")

	// Check count and content
	count := len(set.data)
    AssertEqInt(t, count, 3, "Count wrong")
    AssertTrue(t, set.Contains("B"), "Contents of set")
    AssertFalse(t, set.Contains("Patagonia"), "Contents of set")

	// Check conversion to slice
	size := len(set.AsSlice())
    AssertEqInt(t, size, 3, "Size of slice")

	// Check overwrite from slice
	set = NewSetOfString()
	set.Add("original")
	set.Overwrite([]string{"A", "B", "C"})
	AssertEqInt(t, len(set.data), 3, "length of set")
	AssertTrue(t, set.Contains("B"), "set contents")
}
