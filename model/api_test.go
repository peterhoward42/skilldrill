package model

import (
	"github.com/peterhoward42/skilldrill/util/testutil"
	"testing"
)

//-----------------------------------------------------------------------------
// The basics - smoke tests.
//-----------------------------------------------------------------------------

func TestBasics(t *testing.T) {
	// This exercises the core set of model creation and addition apis, avoiding
	// error conditions. We do not inspect the model built here, because we
	// prefer to kill two birds with one stone and avoid duplicating that logic,
	// by delegating that to a separate serialization test. (See
	// serialize_test.go)
	buildSimpleModel(t)
}

//-----------------------------------------------------------------------------
// Adding things - delibarately stimulating errors
//-----------------------------------------------------------------------------

func TestAddPersonDuplicate(t *testing.T) {
	api := buildSimpleModel(t)
	err := api.AddPerson("fred.bloggs")
	testutil.AssertErrGenerated(t, err, "already exists", "Build simple model.")
}

func TestAddSkillUnknownParent(t *testing.T) {
	api := buildSimpleModel(t)
	_, err := api.AddSkill(SKILL, "title", "desc", 99999)
	testutil.AssertErrGenerated(t, err, "Unknown parent",
		"Adding skill to unknown parent")
}

func TestAddSkillToNonCategory(t *testing.T) {
	api := NewApi()
	rootUid, _ := api.AddSkill(SKILL, "", "", 99999)
	_, err := api.AddSkill(SKILL, "", "", rootUid)
	testutil.AssertErrGenerated(t, err, "must be a category",
		"Adding skill to non-category")
}

//-----------------------------------------------------------------------------
// Give a person a skill - delibarately stimulating errors
//-----------------------------------------------------------------------------

func TestBestowSkillToSpuriousPerson(t *testing.T) {
	api := NewApi()
	skill, _ := api.AddSkill(SKILL, "", "", -1)
	err := api.GivePersonSkill("nosuch.person", skill)
	testutil.AssertErrGenerated(t, err, "Person does not exist",
		"Bestow skill to unknown person")
}

func TestBestowSpuriousSkillToPerson(t *testing.T) {
	api := NewApi()
	api.AddPerson("fred.bloggs")
	err := api.GivePersonSkill("fred.bloggs", 9999)
	testutil.AssertErrGenerated(t, err, "Skill does not exist",
		"Should object to no such skill")
}

func TestBestowCategorySkill(t *testing.T) {
	api := NewApi()
	skill, _ := api.AddSkill(CATEGORY, "", "", -1)
	api.AddPerson("fred.bloggs")
	err := api.GivePersonSkill("fred.bloggs", skill)
	testutil.AssertErrGenerated(t, err, "Cannot give someone a CATEGORY skill",
		"Give someone a category not a skill")
}

func TestEmailsAreLowerCased(t *testing.T) {
	api := NewApi()
	skill, _ := api.AddSkill(SKILL, "", "", -1)
	api.AddPerson("fred.bloggs")
    // Note email address differs with upper case to that used to register
    // the person.
	err := api.GivePersonSkill("fred.Bloggs", skill)
	testutil.AssertNilErr(t, err, "Using uppercase in email.")
}

//-----------------------------------------------------------------------------
// Helper functions
//-----------------------------------------------------------------------------

func buildSimpleModel(t *testing.T) *Api {
	api := NewApi()
	api.AddPerson("fred.bloggs")
	api.AddPerson("john.Smith") // deliberate inclusion of upper case letter
	rootId, _ := api.AddSkill(CATEGORY, "root title",
		"root description", -1)
	skillA, _ := api.AddSkill(
		CATEGORY, "A title", "child A description", rootId)
	skillB, _ := api.AddSkill(
		CATEGORY, "B title", "child B description", rootId)
	skillC, _ := api.AddSkill(
		SKILL, "grandchild", "description", skillA)
	err := api.GivePersonSkill("fred.bloggs", skillC)
	testutil.AssertNilErr(t, err, "Give person skill error")

	_ = skillB

	return api
}
