package model

import (
	"strings"
	"testing"
)


//-----------------------------------------------------------------------------
// Exercise serialization
//-----------------------------------------------------------------------------

func TestSerialize(t *testing.T) {
	api := buildSimpleModel(t)
	serialized, err := api.Serialize()
	if err != nil {
		t.Errorf("serialize failed: %v", err.Error())
		return
	}
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
			t.Errorf("This string missing from serialized: %s", fragment)
		}
	}
}

func TestDeSerialize(t *testing.T) {
	orig := buildSimpleModel(t)
	serialized, err := orig.Serialize()
	if err != nil {
		t.Errorf("serialize failed: %v", err.Error())
		return
	}
	api, err := NewFromSerialized(serialized)

    // Recovered the version ok?
    if q := api.SerializeVers; q != 1 {
		t.Errorf("Serialization version wrong: %d, expected 1", q)
	}

    // Recovered the skills ok?
    if q := api.Skills; len(q) != 4 {
		t.Errorf("Skill list size wrong: %d, expected 4", len(q))
	}
    if skill := api.Skills[3]; skill.Uid != 4 {
		t.Errorf("Skill uid wrong: %d, expected 4", skill.Uid)
	}



	//fmt.Printf("\nDefault repr of restored: <%v>\n", restored)
}
