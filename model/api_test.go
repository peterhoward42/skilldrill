package model

import (
	"strings"
	"testing"
)

//-----------------------------------------------------------------------------
// Adding things without stimulating errors
//-----------------------------------------------------------------------------

func TestAdditions(t *testing.T) {
	buildSimpleModel(t)
}

//-----------------------------------------------------------------------------
// Adding things - delibarately stimulating errors
//-----------------------------------------------------------------------------

func TestAddPersonDuplicate(t *testing.T) {
	api := buildSimpleModel(t)
	err := api.AddPerson("fred.bloggs")
	if err == nil {
		t.Errorf("Should have objected to duplicated addition of fred.")
		return
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Errorf("Error message looks wrong")
	}
}

func TestAddSkillUnknownParent(t *testing.T) {
	api := buildSimpleModel(t)

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
// Sanity test - model building done right
//-----------------------------------------------------------------------------

func TestModelContent(t *testing.T) {
	api := buildSimpleModel(t)
	if len(api.Skills) != 4 {
		t.Errorf("Should be 4 skills")
		return
	}
	if len(api.People) != 2 {
		t.Errorf("Should be 2 people")
		return
	}
	if api.SkillHoldings == nil {
		t.Errorf("Skill holdings ptr is not initialised.")
		return
	}
	mapSiz := len(api.SkillHoldings.SkillsOfPerson)
	if mapSiz != 1 {
		t.Errorf("SkillsOfPeople map should have 1 key, but has: %d", mapSiz)
		return
	}
}

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

    // Recovered the version ok?
    if q := api.SerializeVers; q != 1 {
		t.Errorf("Serialization version wrong: %d, expected 1", q)
	}

    // Recovered the skills ok?
    if q := api.Skills; len(q) != 4 {
		t.Errorf("Skill list size wrong: %d, expected 4", q)
	}
    if skill := api.Skills[3]; skill.Uid != 4 {
		t.Errorf("Skill uid wrong: %d, expected 4", skill.Uid)
	}



	//fmt.Printf("\nDefault repr of restored: <%v>\n", restored)
}

//-----------------------------------------------------------------------------
// Helper functions
//-----------------------------------------------------------------------------

func buildSimpleModel(t *testing.T) *Api {
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
	err := api.GivePersonSkill("fred.bloggs", skillC)
	if err != nil {
		t.Errorf("GivePersonSkill() failed: %v", err.Error())
	}

	_ = skillB

	return api
}
