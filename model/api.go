package model

import (
	"errors"
)

type Api struct {
	treeRoot      *treeNode
	people        map[string]person // keyed on email
	skillHoldings *skillHoldings
}

func NewApi() *Api {
	return &Api{
		people:        map[string]person{},
		skillHoldings: newSkillHoldings(),
	}
}

// Person related API functions

func (api *Api) AddPerson(email string) error {
	// disallow duplicate additions
	_, existingPerson := api.people[email]
	if existingPerson {
		return errors.New("person already exists")
	}
	api.people[email] = person{email}
	return nil
}

func (api *Api) PersonIsKnown(email string) bool {
	_, ok := api.people[email]
	return ok
}

// Skill related API functions

func (api *Api) AddSkill(title string, desc string, parent *treeNode) {
    // When the skill tree root is yet to be initialised, we use this
    // incoming skill to create one, and overwrite the new node's parent field to
    // be nil
    if api.treeRoot == nil {
        api.treeRoot = newTreeNode(title, desc, nil)
        return
    }
    parent.addChild(newTreeNode(title, desc, parent))
}
