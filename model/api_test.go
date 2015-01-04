package model

import (
	"fmt"
	"strings"
	"testing"
)

//-----------------------------------------------------------------------------
// Adding things without stimulating errors
//-----------------------------------------------------------------------------

func TestAdditions(t *testing.T) {
	buildSimpleModel()
}

//-----------------------------------------------------------------------------
// Adding things - delibarately stimulating errors
//-----------------------------------------------------------------------------

func TestAddPersonDuplicate(t *testing.T) {
	api := buildSimpleModel()
	_, err := api.AddPerson("fred.bloggs")
	if err == nil {
		t.Errorf("Should have objected to duplicated addition of fred.")
		return
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Errorf("Error message looks wrong")
	}
}

func TestAddSkillUnknownParent(t *testing.T) {
	api := buildSimpleModel()

	// unknown parent
	_, err := api.AddSkill(SKILL, "title", "desc", 99999)
	if err == nil {
		t.Errorf("Should have objected to unknown parent")
		return
	}
	if !strings.Contains(err.Error(), "Unknown parent") {
		t.Errorf("Error message looks wrong")
	}
}

func TestAddSkillToNonCategory(t *testing.T) {
	api := NewApi()
	rootUid, _ := api.AddSkill(SKILL, "", "", 99999)
	_, err := api.AddSkill(SKILL, "", "", rootUid)
	if err == nil {
		t.Errorf("Should have objected to parent not being category")
		return
	}
	if !strings.Contains(err.Error(), "must be a category") {
		t.Errorf("Error message looks wrong")
	}
}

//-----------------------------------------------------------------------------
// Give a person a skill - delibarately stimulating errors
//-----------------------------------------------------------------------------

func TestBestowSkillToSpuriousPerson(t *testing.T) {
	api := NewApi()
	skill, _ := api.AddSkill(SKILL, "", "", -1)
	err := api.GivePersonSkill("nosuch.person", skill)
	if err == nil {
		t.Errorf("Should have objected to unknown person.")
		return
	}
	if !strings.Contains(err.Error(), "Person does not exist") {
		t.Errorf("Error message looks wrong")
	}
}

func TestBestowSpuriousSkillToPerson(t *testing.T) {
	api := NewApi()
	api.AddPerson("fred.bloggs")
	err := api.GivePersonSkill("fred.bloggs", 9999)
	if err == nil {
		t.Errorf("Should have objected to unknown skill.")
		return
	}
	if !strings.Contains(err.Error(), "Skill does not exist") {
		t.Errorf("Error message looks wrong")
	}
}

func TestBestowCategorySkill(t *testing.T) {
	api := NewApi()
	skill, _ := api.AddSkill(CATEGORY, "", "", -1)
	api.AddPerson("fred.bloggs")
	err := api.GivePersonSkill("fred.bloggs", skill)
	if err == nil {
		t.Errorf("Should have objected to skill being category.")
		return
	}
	if !strings.Contains(err.Error(), "Cannot give someone a CATEGORY skill") {
		t.Errorf("Error message looks wrong")
	}
}

//-----------------------------------------------------------------------------
// Exercise serialization
//-----------------------------------------------------------------------------

func TestSerialize(t *testing.T) {
	api := buildSimpleModel()
	b, err := api.Serialize()
	if err != nil {
		t.Errorf("Serialize failed: %v", err.Error())
		return
	}
	fmt.Printf("Result returned from api serialize follows\n%s", string(b))
}

//-----------------------------------------------------------------------------
// Helper functions
//-----------------------------------------------------------------------------

func buildSimpleModel() *Api {
	api := NewApi()
	api.AddPerson("fred.bloggs")
	api.AddPerson("john.smith")
	rootId, _ := api.AddSkill(CATEGORY, "root title",
		"root description", -1)
	skillA, _ := api.AddSkill(
		CATEGORY, "A title", "child A description", rootId)
	skillB, _ := api.AddSkill(
		CATEGORY, "B title", "child B description", rootId)
	skillC, _ := api.AddSkill(
		SKILL, "grandchild", "description", skillA)
	api.GivePersonSkill("fred.bloggs", skillA)

	_ = skillB
	_ = skillC

	return api
}
