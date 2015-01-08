package model

import (
	"fmt"
	"github.com/peterhoward42/skilldrill/util"
	"strings"
	"testing"
)

//-----------------------------------------------------------------------------
// Exercise serialization
//-----------------------------------------------------------------------------

func TestSerialize(t *testing.T) {
	api := buildSimpleModel(t)
	serialized, err := api.Serialize()
	util.AssertNilErr(t, err, "Serialize error")

	fmt.Printf("\nFor reference, serialized content was:...\n%s\n",
		string(serialized))

	// Perform a sample of smoke tests on content returned.
	fragments := []string{
		"skills:",
		"- uid: 1",
		"title: A title",
		"children:",
		"- 4",
		"children: []",
		"people:",
		"- email: fred.bloggs",
		"skillroot: 1",
		"skillholdings:",
		"skillsofperson:",
		"fred.bloggs:",
		"- 4",
		"nextskill: 5",
	}
	got := string(serialized)
	for _, fragment := range fragments {
		if !strings.Contains(got, fragment) {
			util.AssertStrContains(t, got, fragment, "Serialized content")
		}
	}
}

func TestDeSerialize(t *testing.T) {
	orig := buildSimpleModel(t)
	serialized, err := orig.Serialize()
	util.AssertNilErr(t, err, "Serialize error")

	api, err := NewFromSerialized(serialized)
	util.AssertNilErr(t, err, "DeSerialize error")

	// version ?
	util.AssertEqInt32(t, api.SerializeVers, 1, "Serialize version")
}
