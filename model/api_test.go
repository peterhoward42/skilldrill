package model

import (
	"github.com/peterhoward42/skilldrill/util/testutil"
	"testing"
)

//-----------------------------------------------------------------------------
// The basics - smoke tests.
//-----------------------------------------------------------------------------

func TestTrivial(t *testing.T) {
	api, skillIds := buildSimpleModel(t)
	//skillA := skillIds[0]
	//skillAA := skillIds[1]
	skillAB := skillIds[2]
	skillAAA := skillIds[3]



	testutil.AssertTrue(t, api.PersonExists("fred.bloggs"), "Person exists")
	testutil.AssertEqInt(t, len(skillIds), 4, "Number of skills")
	testutil.AssertTrue(t, api.SkillExists(skillAB), "Skill exists")
	title, _ := api.TitleOfSkill(skillAB)
	testutil.AssertEqString(t, title, "AB", "Title is right")
    hasSkill, _ := api.PersonHasSkill(skillAAA, "fred.bloggs")
	testutil.AssertTrue(t, hasSkill, "Person has skill")
    hasSkill, _ = api.PersonHasSkill(skillAAA, "john.smith")
	testutil.AssertFalse(t, hasSkill, "Person has skill")
}

//-----------------------------------------------------------------------------
// Helper functions
//-----------------------------------------------------------------------------

func buildSimpleModel(t *testing.T) (api *Api, skillIds []int) {
	// Don't change this ! - many tests are dependent on its behaviour and the
	// UIDs generated for the skills added.
	api = NewApi()
	api.AddPerson("fred.bloggs")
	api.AddPerson("john.smith")
	skillA, _ := api.AddSkillNode("A title", "A description", -1)
	skillAA, _ := api.AddSkillNode("AA", "AA description", skillA)
	skillAB, _ := api.AddSkillNode("AB", "AB description", skillA)
	skillAAA, _ := api.AddSkillNode("AAA", "AAA description", skillAA)
	api.GivePersonSkill("fred.bloggs", skillAAA)

	api.ToggleSkillCollapsed("fred.bloggs", skillAA)

	//              A(1)
	//        AA(2)      AB(3)
	// AAA(4)

	_ = skillAB

	return api, []int{skillA, skillAA, skillAB, skillAAA}
}
