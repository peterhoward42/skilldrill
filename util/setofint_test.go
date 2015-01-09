package util

import (
	"testing"
)

func TestSetOfInt(t *testing.T) {
	set := NewSetOfInt()

	// Do some adding, including duplicate addition
	set.Add(1)
	set.Add(2)
	set.Add(3) // note is a duplicate add
	set.Add(3)

	// Check count and content
	count := len(set.data)
    AssertEqInt(t, count, 3, "Count wrong")
    AssertTrue(t, set.Contains(2), "Contents of set")
    AssertFalse(t, set.Contains(999), "Contents of set")

	// Check conversion to slice
	size := len(set.AsSlice())
    AssertEqInt(t, size, 3, "Size of slice")

	// Check overwrite from slice
	set = NewSetOfInt()
	set.Add(99)
	set.Overwrite([]int{1, 2, 3})
	AssertEqInt(t, len(set.data), 3, "length of set")
	AssertTrue(t, set.Contains(1), "set contents")
}
