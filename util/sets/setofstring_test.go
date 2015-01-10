package sets

import (
	"github.com/peterhoward42/skilldrill/util/testutil"
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
	testutil.AssertEqInt(t, count, 3, "Count wrong")
	testutil.AssertTrue(t, set.Contains("B"), "Contents of set")
	testutil.AssertFalse(t, set.Contains("Patagonia"), "Contents of set")

	// Check conversion to slice
	size := len(set.AsSlice())
	testutil.AssertEqInt(t, size, 3, "Size of slice")

	// Check overwrite from slice
	set = NewSetOfString()
	set.Add("original")
	set.Overwrite([]string{"A", "B", "C"})
	testutil.AssertEqInt(t, len(set.data), 3, "length of set")
	testutil.AssertTrue(t, set.Contains("B"), "set contents")
}