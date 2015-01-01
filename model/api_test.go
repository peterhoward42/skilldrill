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
    api.AddSkill("root title", "root description", nil)
	return api
}
