package model

import (
	"errors"
)

//---------------------------------------------------------------------------
// The fundamental type
//---------------------------------------------------------------------------

type Api struct {
	skillRoot     *skillNode         // root of taxonomy tree
	people        map[string]*person // keyed on email
	skillHoldings *skillHoldings     // who has what skill?
	skillFromId   map[int64]*skillNode
	nextUid       int64
}

func NewApi() *Api {
	return &Api{
		people:        map[string]*person{},
		skillHoldings: newSkillHoldings(),
		skillFromId:   map[int64]*skillNode{},
	}
}

//---------------------------------------------------------------------------
// API About people
//---------------------------------------------------------------------------

func (api *Api) AddPerson(email string) (uid int64, err error) {
	// disallow duplicate additions
	_, existingPerson := api.people[email]
	if existingPerson {
		return -1, errors.New("person already exists")
	}
	uid = api.makeUid()
	api.people[email] = newPerson(uid, email)
	return uid, nil
}

func (api *Api) PersonIsKnown(email string) bool {
	_, ok := api.people[email]
	return ok
}

//---------------------------------------------------------------------------
// API About skills
//---------------------------------------------------------------------------

// The role parameter should be one of the constants SKILL or CATEGORY.
// When the skill tree is empty, this skill will be added as the root, and
// the parentUid parameter is ignored.
// If you attempt to add children to a node that is not a CATEGORY, an error is
// produced.
func (api *Api) AddSkill(role int, title string, desc string,
	parentUid int64) (uid int64, err error) {

	// Special case when tree is empty
	if api.skillRoot == nil {
		uid = api.makeUid()
		newNode := newSkillNode(uid, role, title, desc, nil)
		api.skillFromId[uid] = newNode
		api.skillRoot = newNode
		return
	}
	parentNode, ok := api.skillFromId[parentUid]
	if !ok {
		err = errors.New("Unknown parent.")
		return
	}
	if parentNode.role != CATEGORY {
		err = errors.New("Parent must be a category node")
		return
	}
	newNode := newSkillNode(uid, role, title, desc, parentNode)
	api.skillFromId[uid] = newNode
	parentNode.addChild(newNode)
	return
}

//---------------------------------------------------------------------------
// API About Which Skills People Have
//---------------------------------------------------------------------------

func (api *Api) GivePersonSkill(email string, skillId int64) error {
	foundPerson, ok := api.people[email]
	if !ok {
		return errors.New("person does not exist")
	}
	skill := api.skillFromId[skillId]
	api.skillHoldings.bind(skill, foundPerson)
	return nil
}

//---------------------------------------------------------------------------
// Not exported
//---------------------------------------------------------------------------

func (api *Api) makeUid() int64 {
	api.nextUid++
	return api.nextUid
}
