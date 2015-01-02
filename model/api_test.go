package model

import (
	"strings"
	"testing"
)

func TestAdditions(t *testing.T) {
	buildSimpleModel()
}

func TestErrors(t *testing.T) {
	api := buildSimpleModel()
	_, err := api.AddPerson("fred.bloggs")
	if err == nil {
		t.Errorf("Should have objected to duplicated addition of fred.")
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Errorf("Error message looks wrong")
	}
}

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
