package model

import (
	"github.com/peterhoward42/skilldrill/util/testutil"
	"testing"
)

//-----------------------------------------------------------------------------
// The basics - smoke tests.
//-----------------------------------------------------------------------------

func TestBasics(t *testing.T) {
	api, skillIds := buildSimpleModel(t)
    _ = api
    _ = skillIds
}

func TestSimpleContent(t *testing.T) {
	api, skillIds := buildSimpleModel(t)
    testutil.AssertTrue(t, api.model.holdings.personExists(
                "fred.bloggs"), "Content correct.")
    skillIds, depths, fredHas := api.EnumerateTree("fred.bloggs")
    testutil.AssertEqInt(t, len(skillIds), 3, "tree data")

    _ = skillIds
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
