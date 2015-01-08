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
	if count != 3 {
		t.Errorf("Count wrong %d, expected %d", count, 3)
	}
	if !set.Contains("B") {
		t.Errorf("Should say contains 2 is true")
	}
	if set.Contains("garbage") {
		t.Errorf(`Should say does not contain "garbage"`)
	}

	// Check conversion to slice
	asSlice := set.AsSlice()
	count = len(asSlice)
	if count != 3 {
		t.Errorf("Slice length wrong: %d, expected 3.", count)
	}

	// Check overwrite from slice
	set = NewSetOfString()
	set.Add("furniture")
	set.Overwrite([]string{"apple", "pear", "orange"})
	AssertEqInt(t, len(set.data), 3, "length of set")
	AssertTrue(t, set.Contains("pear"), "set contents")
}
