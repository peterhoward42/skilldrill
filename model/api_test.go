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
	testutil.AssertErrGenerated(t, err, PersonExists, "Build simple model.")
}

func TestAddSkillUnknownParent(t *testing.T) {
	api := buildSimpleModel(t)
	_, err := api.AddSkill(Skill, "title", "desc", 99999)
	testutil.AssertErrGenerated(t, err, UnknownParent,
		"Adding skill to unknown parent")
}

func TestAddSkillToNonCategory(t *testing.T) {
	api := NewApi()
	rootUid, _ := api.AddSkill(Skill, "", "", 99999)
	_, err := api.AddSkill(Skill, "", "", rootUid)
	testutil.AssertErrGenerated(t, err, ParentNotCategory,
		"Adding skill to non-category")
}

//-----------------------------------------------------------------------------
// Give a person a skill - delibarately stimulating errors
//-----------------------------------------------------------------------------

func TestBestowSkillToSpuriousPerson(t *testing.T) {
	api := NewApi()
	skill, _ := api.AddSkill(Skill, "", "", -1)
	err := api.GivePersonSkill("nosuch.person", skill)
	testutil.AssertErrGenerated(t, err, UnknownPerson,
		"Bestow skill to unknown person")
}

func TestBestowSpuriousSkillToPerson(t *testing.T) {
	api := NewApi()
	api.AddPerson("fred.bloggs")
	err := api.GivePersonSkill("fred.bloggs", 9999)
	testutil.AssertErrGenerated(t, err, UnknownSkill,
		"Should object to no such skill")
}

func TestBestowCategorySkill(t *testing.T) {
	api := NewApi()
	skill, _ := api.AddSkill(Category, "", "", -1)
	api.AddPerson("fred.bloggs")
	err := api.GivePersonSkill("fred.bloggs", skill)
	testutil.AssertErrGenerated(t, err, CategoryDisallowed,
		"Give someone a category not a skill")
}

func TestEmailsAreLowerCased(t *testing.T) {
	api := NewApi()
	skill, _ := api.AddSkill(Skill, "", "", -1)
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
	skillA, _ := api.AddSkill(Category, "A title", "A description", -1)
	skillAA, _ := api.AddSkill(Category, "AA", "AA description", skillA)
	skillAB, _ := api.AddSkill(Category, "AB", "AB description", skillA)
	skillAAA, _ := api.AddSkill(Skill, "AAA", "AAA description", skillAA)
	api.GivePersonSkill("fred.bloggs", skillAAA)

	err := api.CollapseSkill("fred.bloggs", skillAA)
	testutil.AssertNilErr(t, err, "CollapseSkill during dev only")

	_ = skillAB

	return api
}
