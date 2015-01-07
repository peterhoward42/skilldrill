package util

import (
	"testing"
)

func TestSetOfInt(t *testing.T) {
	set := NewSetOfInt()
	set.Add(1)
	set.Add(2)
	set.Add(3) // note is a duplicate add
	set.Add(3)

	count := len(set.data)
	if count != 3 {
		t.Errorf("Count wrong %d, expected %d", count, 3)
	}
	if !set.Contains(2) {
		t.Errorf("Should say contains 2 is true")
	}
	if set.Contains(9999) {
		t.Errorf("Should say does not contain 9999")
	}
	asSlice := set.AsSlice()
	count = len(asSlice)
	if count != 3 {
		t.Errorf("Slice length wrong: %d, expected 3.", count)
	}
}
