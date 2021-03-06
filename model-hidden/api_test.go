package model

import (
	"github.com/peterhoward42/skilldrill/util/testutil"
	"sort"
	"strings"
	"testing"
)

//-----------------------------------------------------------------------------
// The basics - smoke tests.
//-----------------------------------------------------------------------------

/*
This exercises the core set of model creation and addition apis, avoiding error
conditions. We do not inspect the model built here, because we prefer to kill
two birds with one stone and avoid duplicating that logic, by delegating that
to a separate serialization test. (See serialize_test.go)
*/
func TestBasics(t *testing.T) {
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
// Adding things - checking interventions in the Api layer
//-----------------------------------------------------------------------------

func TestChildrenOrderedAlphabetically(t *testing.T) {
	// Ensure the children of a parent in common, are kept in alphabetical
	// order when they are added (deliberately) no so.
	api := buildSimpleModel(t)
	childIds := api.skillFromId[1].Children
	titles := []string{}
	for _, child := range childIds {
		titles = append(titles, api.skillFromId[child].Title)
	}
	testutil.AssertTrue(t, sort.StringsAreSorted(titles),
		"Children are not sorted.")
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
	testutil.AssertErrGenerated(t, err, CannotBestowCategory,
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
// Editing operations with and without errors
//-----------------------------------------------------------------------------

func TestSkillEditsErrors(t *testing.T) {
	api := NewApi()
	skill, err := api.AddSkill(Skill, "Orig Title", "Orig desc.", -1)
	testutil.AssertNilErr(t, err, "Adding skill")

	err = api.SetSkillTitle(skill, "New Title")
	testutil.AssertNilErr(t, err, "Setting skill title.")
	testutil.AssertEqString(t, api.skillFromId[skill].Title, "New Title",
		"Setting skill title")

	err = api.SetSkillTitle(999, "who cares")
	testutil.AssertErrGenerated(t, err, UnknownSkill, "Set skill title.")
	err = api.SetSkillTitle(skill, strings.Repeat("X", 40))
	testutil.AssertErrGenerated(t, err, TooLong, "Setting skill title.")

	err = api.SetSkillDesc(skill, "New Desc")
	testutil.AssertNilErr(t, err, "Setting skill desc.")
	testutil.AssertEqString(t, api.skillFromId[skill].Desc, "New Desc",
		"Setting skill desc")

	err = api.SetSkillDesc(999, "New Desc")
	testutil.AssertErrGenerated(t, err, UnknownSkill, "Set skill desc.")
	err = api.SetSkillDesc(skill, strings.Repeat("X", 500))
	testutil.AssertErrGenerated(t, err, TooLong, "Setting skill desc.")
}

func TestMoveSkillInTree(t *testing.T) {
	api := buildSimpleModel(t)
	err := api.ReParentSkill(4, 1)
	testutil.AssertNilErr(t, err, "Re parenting skill")
	testutil.AssertEqSliceInt(t, api.skillFromId[1].Children,
		[]int{3, 4, 2}, "Re parenting skill")
	testutil.AssertEqInt(t, api.skillFromId[4].Parent, 1,
		"Re parenting skill")

	err = api.ReParentSkill(4, 999)
	testutil.AssertErrGenerated(t, err, UnknownSkill, "Re parenting skill")
	err = api.ReParentSkill(999, 1)
	testutil.AssertErrGenerated(t, err, UnknownSkill, "Re parenting skill")

	err = api.ReParentSkill(1, 4)
	testutil.AssertErrGenerated(t, err, IllegalWithRoot, "Re parenting skill")

	api = buildSimpleModel(t)
	err = api.ReParentSkill(3, 4)
	testutil.AssertErrGenerated(t, err, ParentNotCategory, "Re parenting skill")
}

func TestRemovePerson(t *testing.T) {
	api := buildSimpleModel(t)
	err := api.RemovePerson("fred.bloggs")
	testutil.AssertNilErr(t, err, "Remove Person")
	testutil.AssertFalse(t, api.PersonExists("fred.bloggs"), "Remove Person")
	err = api.RemovePerson("no suchperson")
	testutil.AssertErrGenerated(t, err, UnknownPerson, "Remove Person")
}

func TestRemoveSkill(t *testing.T) {
	api := buildSimpleModel(t)

	// Try to remove a skill with children
	err := api.RemoveSkill(3)
	testutil.AssertErrGenerated(t, err, CannotRemoveSkillWithChildren,
		"Remove Skill")

	// Try to remove an unnkown skill
	err = api.RemoveSkill(999)
	testutil.AssertErrGenerated(t, err, UnknownSkill, "Remove Skill")

	// Try to remove the root skill
	err = api.RemoveSkill(1)
	testutil.AssertErrGenerated(t, err, CannotRemoveRootSkill,
		"Remove Skill")

	// Try to remove a skill that has people registered agains it
	err = api.RemoveSkill(4)
	testutil.AssertErrGenerated(t, err, CannotRemoveSkillHeld,
		"Remove Skill")

	// Before we test for legitimate usage, we will get john smith to collapse
	// the tree node that is about to go, so that we can check that it no
	// longer shows up in his collapsed skill node set.
	err = api.CollapseSkill("john.smith", 2)
	testutil.AssertNilErr(t, err, "Collapse Skill.")
	err = api.RemoveSkill(2)
	testutil.AssertNilErr(t, err, "Remove Skill")

	// Check that the skill that was removed has been forgotten in every way
	_, err = api.PeopleWithSkill(2)
	testutil.AssertErrGenerated(t, err, UnknownSkill, "Remove Skill")
	skills, _, err := api.EnumerateTree("john.smith")
	testutil.AssertEqSliceInt(t, skills, []int{1, 3, 4}, "Remove Skill")
	testutil.AssertFalse(t,
		api.UiStates["john.smith"].CollapsedNodes.Contains(2), "Remove Skill")
}

//-----------------------------------------------------------------------------
// Exercise Queries
//-----------------------------------------------------------------------------

func TestSkillQueries(t *testing.T) {
	api := buildSimpleModel(t)

	// Proper use
	title, desc, descInContext, contextAlone, err := api.SkillWording(4)
	testutil.AssertNilErr(t, err, "Skill wording getter")
	testutil.AssertStrContains(t, title, "AAA", "Skill wording getter")
	testutil.AssertStrContains(t, desc, "AAA desc", "Skill wording getter")
	testutil.AssertStrContains(t, descInContext,
		"A description>>>AA description>>>AAA description",
		"Skill wording getter")
	testutil.AssertStrContains(t, contextAlone,
		"A description>>>AA description", "Skill wording getter")

	// Illegal skill id
	_, _, _, _, err = api.SkillWording(999)
	testutil.AssertErrGenerated(t, err, UnknownSkill, "Skill wording getter")
}

func TestPeopleWithSkillQuery(t *testing.T) {
	api := buildSimpleModel(t)
	emails, err := api.PeopleWithSkill(4)
	testutil.AssertNilErr(t, err, "People with skill getter")
	testutil.AssertEqSliceString(t, emails, []string{"fred.bloggs"},
		"People with skill getter")

	emails, err = api.PeopleWithSkill(999)
	testutil.AssertErrGenerated(t, err, UnknownSkill,
		"People with skill getter")

	emails, err = api.PeopleWithSkill(1)
	testutil.AssertErrGenerated(t, err, CannotBestowCategory,
		"People with skill getter")
}

func TestHasPersonSkillQuery(t *testing.T) {
	api := buildSimpleModel(t)

	// Proper usage
	hasSkill, err := api.PersonHasSkill("fred.bloggs", 4)
	testutil.AssertNilErr(t, err, "Person has skill getter")
	testutil.AssertTrue(t, hasSkill, "Person has skill getter")

	hasSkill, err = api.PersonHasSkill("john.smith", 4)
	testutil.AssertNilErr(t, err, "Person has skill getter")
	testutil.AssertFalse(t, hasSkill, "Person has skill getter")

	// Error generation
	hasSkill, err = api.PersonHasSkill("no such person", 4)
	testutil.AssertErrGenerated(t, err, UnknownPerson,
		"People with skill getter")

	hasSkill, err = api.PersonHasSkill("fred.bloggs", 999)
	testutil.AssertErrGenerated(t, err, UnknownSkill,
		"People with skill getter")

	hasSkill, err = api.PersonHasSkill("fred.bloggs", 1)
	testutil.AssertErrGenerated(t, err, CannotBestowCategory,
		"People with skill getter")
}

func TestEnumerateTree(t *testing.T) {
	api := buildSimpleModel(t)
	skills, depths, err := api.EnumerateTree("fred.bloggs")
	testutil.AssertNilErr(t, err, "Tree enumerator")
	testutil.AssertEqSliceInt(t, skills, []int{1, 3, 2}, "Tree enumerator")
	testutil.AssertEqSliceInt(t, depths, []int{0, 1, 1}, "Tree enumerator")

	skills, depths, err = api.EnumerateTree("nosuch person")
	testutil.AssertErrGenerated(t, err, UnknownPerson, "Tree enumerator")
}

//-----------------------------------------------------------------------------
// Operate virtualized UXP - stimulating errors
//-----------------------------------------------------------------------------

func TestCollapseSkillErrors(t *testing.T) {
	api := buildSimpleModel(t)
	err := api.CollapseSkill("fred.bloggs", 9999)
	testutil.AssertErrGenerated(t, err, UnknownSkill, "Collapse node.")
	err = api.CollapseSkill("nosuchemail", 1)
	testutil.AssertErrGenerated(t, err, UnknownPerson, "Collapse node.")
}

//-----------------------------------------------------------------------------
// Helper functions
//-----------------------------------------------------------------------------

func buildSimpleModel(t *testing.T) *Api {
	// Don't change this ! - many tests are dependent on its behaviour and the
	// UIDs generated for the skills added.
	api := NewApi()
	api.AddPerson("fred.bloggs")
	api.AddPerson("john.Smith") // deliberate inclusion of upper case letter
	skillA, _ := api.AddSkill(Category, "A title", "A description", -1)
	// Note AB and AA are added to a parent in common, in an order that makes
	// their enumeration in the order that they are added, NOT in alphabetical
	// order.
	skillAB, _ := api.AddSkill(Category, "AB", "AB description", skillA)
	skillAA, _ := api.AddSkill(Category, "AA", "AA description", skillA)
	skillAAA, _ := api.AddSkill(Skill, "AAA", "AAA description", skillAA)
	api.GivePersonSkill("fred.bloggs", skillAAA)

	err := api.CollapseSkill("fred.bloggs", skillAA)
	testutil.AssertNilErr(t, err, "CollapseSkill during dev only")

	//              A(1)
	//        AA(3)      AB(2)
	// AAA(4)

	_ = skillAB

	return api
}
