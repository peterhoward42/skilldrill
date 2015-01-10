package model

import (
	"github.com/peterhoward42/skilldrill/util/testutil"
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
	testutil.AssertNilErr(t, err, "Serialize error")

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
			testutil.AssertStrContains(t, got, fragment, "Serialized content")
		}
	}
}

func TestDeSerialize(t *testing.T) {
	orig := buildSimpleModel(t)
	serialized, err := orig.Serialize()
	testutil.AssertNilErr(t, err, "Serialize error")

	// Ensure de-serialize does not generate errors in of itself.
	api, err := NewFromSerialized(serialized)
	testutil.AssertNilErr(t, err, "DeSerialize error")

	// Probe correctness of data...

	testutil.AssertEqInt(t, api.SerializeVers, 1, "Serialize version")
	checkSkills(t, api)
	checkPeople(t, api)
	checkSkillFromId(t, api)
	checkPersFromMail(t, api)
	testutil.AssertEqInt(t, api.SkillRoot, 1, "Skill Root")
	checkSkillHoldings(t, api)
	testutil.AssertEqInt(t, api.NextSkill, 5, "Next skill")
}

func checkSkills(t *testing.T, api *Api) {
	// Right number ?
	n := len(api.Skills)
	testutil.AssertEqInt(t, n, 4, "Number of skills")

	// Content of one of the skills
	skill := api.Skills[0]
	testutil.AssertEqInt(t, skill.Uid, 1, "Skill id.")
	testutil.AssertEqString(t, skill.Role, CATEGORY, "Role.")
	testutil.AssertEqString(t, skill.Title, "root title", "Title.")
	testutil.AssertEqString(t, skill.Desc, "root description", "Description.")
	testutil.AssertEqInt(t, skill.Parent, -1, "Parent")
	testutil.AssertEqSliceInt(t, skill.Children, []int{2, 3}, "Children.")
}

func checkPeople(t *testing.T, api *Api) {
	n := len(api.People)
	testutil.AssertEqInt(t, n, 2, "Number of people")
	john := api.People[1]
    // This also checks that emails added that include uppercase letters,
    // are coerced to lower case by the AddXXX methods.
	testutil.AssertEqString(t, john.Email, "john.smith", "Person name.")
}

func checkSkillFromId(t *testing.T, api *Api) {
	skill := api.skillFromId[1]
	testutil.AssertEqInt(t, skill.Uid, 1, "Skill id.")
	skill = api.skillFromId[2]
	testutil.AssertEqInt(t, skill.Uid, 2, "Skill id.")
}

func checkPersFromMail(t *testing.T, api *Api) {
	pp := api.persFromMail["john.smith"]
	testutil.AssertEqString(t, pp.Email, "john.smith", "Email.")
}

func checkSkillHoldings(t *testing.T, api *Api) {
	skillset := api.SkillHoldings.SkillsOfPerson["fred.bloggs"].AsSlice()
	testutil.AssertEqSliceInt(t, skillset, []int{4}, "Skills of person.")

	people := api.SkillHoldings.PeopleWithSkill[4].AsSlice()
	testutil.AssertEqSliceString(t, people, []string{"fred.bloggs"},
		"People with skill.")
}
