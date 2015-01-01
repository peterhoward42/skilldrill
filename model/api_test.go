package model

import (
	"testing"
)

// Add people and exercise queries
func TestAddPeople(t *testing.T) {
	api := NewApi()

	err := api.AddPerson("fred.bloggs")
	if err != nil {
		t.Errorf("Problem adding person: %v", err)
	}
	err = api.AddPerson("john.smith")
	if api.PersonIsKnown("fred.bloggs") == false {
		t.Errorf("fred.bloggs should exist")
	}
	if api.PersonIsKnown("nosuch.person") == true {
		t.Errorf("nosuch.person should not be found")
	}
	err = api.AddPerson("fred.bloggs")
	if err == nil {
		t.Errorf("adding person that exists should cause error")
	}

	// remove person that exists

	// remove person that does not exist

	// email addresses coerced to lower case when added
}

func TestAddSkills(t *testing.T) {
	//t.Errorf("Not implemented")
}

func TestEditSkillsOfPerson(t *testing.T) {
	//t.Errorf("Not implemented")
}
