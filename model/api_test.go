package model

import (
	"github.com/peterhoward42/skilldrill/util/testutil"
	"testing"
)

//-----------------------------------------------------------------------------
// The basics - smoke tests.
//-----------------------------------------------------------------------------

func TestSimpleErrorFreeUse(t *testing.T) {
	api, skillIds := buildSimpleModel(t)
	testutil.AssertTrue(t, api.PersonExists("fred.bloggs"), "Person exists")
	testutil.AssertEqInt(t, len(skillIds), 4, "Number of skills")
	testutil.AssertTrue(t, api.SkillExists(skillIds["skillAB"]), "Skill exists")
	title, _ := api.TitleOfSkill(skillIds["skillAB"])
	testutil.AssertEqString(t, title, "AB", "Title is right")
	hasSkill, _ := api.PersonHasSkill(skillIds["skillAAA"], "fred.bloggs")
	testutil.AssertTrue(t, hasSkill, "Person has skill")
	hasSkill, _ = api.PersonHasSkill(skillIds["skillAAA"], "john.smith")
	testutil.AssertFalse(t, hasSkill, "Person has skill")
}

func TestEnumerateTree(t *testing.T) {
	api, _ := buildSimpleModel(t)
	skillSeq := api.EnumerateTree("fred.bloggs")
	expected := []string{"A title", "AA", "AB"}
	testutil.AssertEqInt(t, len(skillSeq), len(expected), "TestEnumerateTree")
	for idx, skillId := range skillSeq {
		receivedTitle, _ := api.TitleOfSkill(skillId)
		testutil.AssertEqString(t, receivedTitle, expected[idx],
			"TestEnumerateTree")
	}
}

func TestCollapseNodeToggles(t *testing.T) {
    api, skills := buildSimpleModel(t)
    skillAA := skills["aa"]
    // buildSimpleModel leaves AA in a collapsed state
    api.ToggleSkillCollapsed("fred.bloggs", skillAA)
	isCollapsed, _ := api.IsCollapsed("fred.bloggs", skillAA)
	testutil.AssertFalse(t, isCollapsed, "TestCollapseNode")
}

//-----------------------------------------------------------------------------
// Helper functions
//-----------------------------------------------------------------------------

func buildSimpleModel(t *testing.T) (api *Api, skillIds map[string]int) {
	// Don't change this ! - many tests are dependent on its behaviour and the
	// UIDs generated for the skills added.
	api = NewApi()
	api.AddPerson("fred.bloggs")
	api.AddPerson("john.smith")
	skillIds = map[string]int{}
	skillIds["skillA"], _ = api.AddSkillNode("A title", "A description", -1)
	skillIds["skillAA"], _ = api.AddSkillNode("AA", "AA description",
		skillIds["skillA"])
	skillIds["skillAB"], _ = api.AddSkillNode("AB", "AB description",
		skillIds["skillA"])
	skillIds["skillAAA"], _ = api.AddSkillNode("AAA", "AAA description",
		skillIds["skillAA"])
	api.GivePersonSkill("fred.bloggs", skillIds["skillAAA"])

	api.ToggleSkillCollapsed("fred.bloggs", skillIds["skillAA"])

	//              A(1)
	//        AA(2)      AB(3)
	// AAA(4)

	return api, skillIds
}
