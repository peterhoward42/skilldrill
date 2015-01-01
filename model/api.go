package model

import (
	"errors"
)

type Api struct {
	treeRoot      *TreeNode
	people        map[string]Person // keyed on email
	skillHoldings *SkillHoldings
}

func NewApi() *Api {
	return &Api{
		people:        map[string]Person{},
		skillHoldings: NewSkillHoldings(),
	}
}

// Person related API functions

func (api *Api) AddPerson(email string) error {
	// disallow duplicate additions
	_, person := api.people[email]
	if person {
		return errors.New("person already exists")
	}
	api.people[email] = Person{email}
	return nil
}

func (api *Api) PersonIsKnown(email string) bool {
	_, ok := api.people[email]
	return ok
}
