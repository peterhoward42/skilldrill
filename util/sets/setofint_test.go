package sets

import (
	"github.com/peterhoward42/skilldrill/util/testutil"
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
	testutil.AssertEqInt(t, count, 3, "Count wrong")
	testutil.AssertTrue(t, set.Contains(2), "Contents of set")
	testutil.AssertFalse(t, set.Contains(999), "Contents of set")

	// Check conversion to slice
	size := len(set.AsSlice())
	testutil.AssertEqInt(t, size, 3, "Size of slice")

	// Check overwrite from slice
	set = NewSetOfInt()
	set.Add(99)
	set.Overwrite([]int{1, 2, 3})
	testutil.AssertEqInt(t, len(set.data), 3, "length of set")
	testutil.AssertTrue(t, set.Contains(1), "set contents")

	// Check removal
	set.Remove(2)
	testutil.AssertEqInt(t, len(set.data), 2, "length of set")
	testutil.AssertTrue(t, set.Contains(1), "set contents")
	testutil.AssertFalse(t, set.Contains(2), "set contents")
	testutil.AssertTrue(t, set.Contains(3), "set contents")

	// Check remove-if-present
	set = NewSetOfInt()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.RemoveIfPresent(9999)
	testutil.AssertEqInt(t, len(set.AsSlice()), 3, "Remove if present.")
	set.RemoveIfPresent(2)
	testutil.AssertEqInt(t, len(set.AsSlice()), 2, "Remove if present.")
	testutil.AssertFalse(t, set.Contains(2), "Remove if present")
}
