/*
The skilldrill model package is a multi-file package that can model a hierachical
set of skills and a set of people who hold some of those skills. The api.go file
exposes the Api type, which provides methods for CRUD operations, while the other
files deal with much of the internal workings.
*/
package model

import (
	"errors"
)

// The Api structure is the fundamental type exposed by the skilldrill model
// package, and provides CRUD interfaces.
type Api struct {
	skillRoot     *skillNode         // root of taxonomy tree
	people        map[string]*person // keyed on email
	skillHoldings *skillHoldings     // who has what skill?
	skillFromId   map[int64]*skillNode
	nextUid       int64
}

// The function NewApi() is a (compulsory) constructor for the Api type.
func NewApi() *Api {
	return &Api{
		people:        map[string]*person{},
		skillHoldings: newSkillHoldings(),
		skillFromId:   map[int64]*skillNode{},
	}
}

// The AddPerson() method adds a person to the model in terms of the user name
// part of their email address. It returns the unique identity it has generated
// for the person, and potentially an error value. It is an error to add a person
// that already exists in the model.
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

// The PersonIsKnown() method returns true if the given person is already present
// in the model.
func (api *Api) PersonIsKnown(email string) bool {
	_, ok := api.people[email]
	return ok
}

/*
The AddSkill() method adds a skill into the model's hierachy of skills.  You
specify the skill in terms of description and title strings. These strings should
describe how they additionally qualify their context in the hierachy, and should
not duplicate context information.  You specify the tree location by providing
the Uid of the parent skill, and the new Uid for the added skill is returned.
The role parameter should be one of the constants SKILL or CATEGORY.  When the
skill tree is empty, this skill will be added as the root, and the parentUid
parameter is ignored.  Errors are generated if you attempt to add a skill to a
node that is not a CATEGORY, or if the parent skill you provide is not
recognized.
*/
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

/*
The GivePersonSkill() method adds the given skill into the set of skills the
model holds for that person.  You are only allowed to give people SKILLS, not
CATEGORIES.  An error is generated if either the person or skill given are not
recognized, or you give a person a CATEGORY rather than a SKILL.
*/
func (api *Api) GivePersonSkill(email string, skillId int64) error {
	foundPerson, ok := api.people[email]
	if !ok {
		return errors.New("Person does not exist.")
	}
	foundSkill, ok := api.skillFromId[skillId]
	if !ok {
		return errors.New("Skill does not exist.")
	}
	if foundSkill.role == CATEGORY {
		return errors.New("Cannot give someone a CATEGORY skill.")
	}
	api.skillHoldings.bind(foundSkill, foundPerson)
	return nil
}

// The makeUid() method is a factory for new unique identifiers. They are unique
// only with respect to the instance of the Api object.
func (api *Api) makeUid() int64 {
	api.nextUid++
	return api.nextUid
}
