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

func (api *Api) AddSkill(title string, desc string, parentUid int64) (uid int64) {
	// When the skill tree root is yet to be initialised, we use this
	// incoming skill to create one, and overwrite the new node's parent field to
	// be nil
	uid = api.makeUid()
	var newNode *skillNode
	if api.skillRoot == nil {
		newNode = newSkillNode(uid, title, desc, nil)
		api.skillRoot = newNode
	} else {
		parentNode := api.skillFromId[parentUid]
		newNode = newSkillNode(uid, title, desc, parentNode)
		parentNode.addChild(newNode)
	}
	api.skillFromId[uid] = newNode
	return
}

//---------------------------------------------------------------------------
// Not exported
//---------------------------------------------------------------------------

func (api *Api) makeUid() int64 {
	api.nextUid++
	return api.nextUid
}
