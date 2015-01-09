package model

import (
	"github.com/peterhoward42/skilldrill/util"
	"strings"
	"testing"
)

//-----------------------------------------------------------------------------
// The functions in this module, have two purposes. The first is to ensure that a
// an Api instance can be serialized and then de-serialized, and in so doing,
// produce a model that is the same as the original one. The second purpose is to
// validate the correct operation of the various addition functions in the api.
// This avoids the latter from being repeated elsewhere.
//-----------------------------------------------------------------------------

func TestSerialize(t *testing.T) {
	api := buildSimpleModel(t)
	serialized, err := api.Serialize()
	util.AssertNilErr(t, err, "Serialize error")

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

	// Ensure de-serialize does not generate errors in of itself.
	api, err := NewFromSerialized(serialized)
	util.AssertNilErr(t, err, "DeSerialize error")

	// Probe correctness of data...

	util.AssertEqInt(t, api.SerializeVers, 1, "Serialize version")
	checkSkills(t, api)
	checkPeople(t, api)
	checkSkillFromId(t, api)
	checkPersFromMail(t, api)
	util.AssertEqInt(t, api.SkillRoot, 1, "Skill Root")
	checkSkillHoldings(t, api)
	util.AssertEqInt(t, api.NextSkill, 5, "Next skill")
}

func checkSkills(t *testing.T, api *Api) {
	// Right number ?
	n := len(api.Skills)
	util.AssertEqInt(t, n, 4, "Number of skills")

	// Content of one of the skills
	skill := api.Skills[0]
	util.AssertEqInt(t, skill.Uid, 1, "Skill id.")
	util.AssertEqString(t, skill.Role, CATEGORY, "Role.")
	util.AssertEqString(t, skill.Title, "root title", "Title.")
	util.AssertEqString(t, skill.Desc, "root description", "Description.")
	util.AssertEqInt(t, skill.Parent, -1, "Parent")
	util.AssertEqSliceInt(t, skill.Children, []int{2, 3}, "Children.")
}

func checkPeople(t *testing.T, api *Api) {
	n := len(api.People)
	util.AssertEqInt(t, n, 2, "Number of people")
	john := api.People[1]
	util.AssertEqString(t, john.Email, "john.smith", "Person name.")
}

func checkSkillFromId(t *testing.T, api *Api) {
	skill := api.skillFromId[1]
	util.AssertEqInt(t, skill.Uid, 1, "Skill id.")
	skill = api.skillFromId[2]
	util.AssertEqInt(t, skill.Uid, 2, "Skill id.")
}

func checkPersFromMail(t *testing.T, api *Api) {
	pp := api.persFromMail["john.smith"]
	util.AssertEqString(t, pp.Email, "john.smith", "Email.")
}

func checkSkillHoldings(t *testing.T, api *Api) {
	skillset := api.SkillHoldings.SkillsOfPerson["fred.bloggs"].AsSlice()
	util.AssertEqSliceInt(t, skillset, []int{4}, "Skills of person.")

	people := api.SkillHoldings.PeopleWithSkill[4].AsSlice()
	util.AssertEqSliceString(t, people, []string{"fred.bloggs"},
		"People with skill.")
}
