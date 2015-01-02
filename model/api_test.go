package model

import (
	"testing"
)

func TestAdditions(t *testing.T) {
	buildSimpleModel()
}

func TestQueries(t *testing.T) {
	api := buildSimpleModel()
	if api.PersonIsKnown("fred.bloggs") == false {
		t.Errorf("fred bloggs is missing")
	}
	if api.PersonIsKnown("garbage") == true {
		t.Errorf("person known false positive")
	}
}

func buildSimpleModel() *Api {
	api := NewApi()
	api.AddPerson("fred.bloggs")
	api.AddPerson("john.smith")
	var unusedUid int64 = -1 // because we are adding the root node
	rootId := api.AddSkill("root title", "root description", unusedUid)
	skillA := api.AddSkill("child A title", "child A description", rootId)
	skillB := api.AddSkill("child B title", "child B description", rootId)
	skillC := api.AddSkill("grandchild", "description", skillA)
	api.GivePersonSkill("fred.bloggs", skillA)

	_ = skillB
	_ = skillC

	return api
}
