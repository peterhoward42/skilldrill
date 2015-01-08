package model

import (
    "runtime"
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

    versionOK(t, api)
    skillListOK(t, api)
    sampleSkillOK(t, api)
}

func versionOK(t *testing.T, api *Api) {
    if q := api.SerializeVers; q != 1 {
		t.Errorf("Serialization version wrong: %d, expected 1", q)
	}
}

func skillListOK(t *testing.T, api *Api) {
    if q := api.Skills; len(q) != 4 {
		t.Errorf("Skill list size wrong: %d, expected 4", len(q))
	}
}

func sampleSkillOK(t *testing.T, api *Api) {
    skill := api.Skills[3];
    foo(t, skill.Uid, 5, "Uid")
}

func foo(t *testing.T, got int32, expected int32, isWrong string) {
    if got != expected {
        var buf = make([]byte, 10000) // has to be big enough
        written := runtime.Stack(buf, false)
		t.Errorf("%s is wrong: %v, expected: %v", isWrong, got, expected)
        t.Error(string(buf[0:written]))
	}
}

