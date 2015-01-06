/*
The skilldrill model package is a multi-file package that can model a hierachical
set of skills and a set of people who hold some of those skills. The api.go file
exposes the Api type, which provides methods for CRUD operations, while the other
files deal with much of the internal workings.
*/
package model

import (
	"errors"
	"gopkg.in/yaml.v2"
)

// The Api structure is the fundamental type exposed by the skilldrill model
// package, and provides CRUD interfaces to do things like adding skills or
// people into the model and registering a person as having a particular skill.
// All model editing operations should be done via Api calls rather than
// accessing the internal objects directly, so that the integrity of various
// supplemental look up tables is preserved. The design intent is that none of
// the Api fields are exported, but the reason that some are, is solely to
// facilitate automated serialization by yaml.Marshal().
type Api struct {
	Skills        []*skillNode
	People        []*person
	skillFromId   map[int32]*skillNode
	persFromMail  map[string]*person
	SkillRoot     int32          // root of taxonomy tree (skill.uid)
	SkillHoldings *skillHoldings // who has what skill?
	nextSkill     int32
}

// The function NewApi() is a (compulsory) constructor for the Api type.
func NewApi() *Api {
	return &Api{
		Skills:        make([]*skillNode, 0),
		People:        make([]*person, 0),
		skillFromId:   make(map[int32]*skillNode),
		persFromMail:  make(map[string]*person),
		SkillRoot:     -1,
		SkillHoldings: newSkillHoldings(),
		nextSkill:     1,
	}
}

// The AddPerson() method adds a person to the model in terms of the user name
// part of their email address. It is an error to add a person that already
// exists in the model.
func (api *Api) AddPerson(email string) (err error) {
	// disallow duplicate additions
	_, existingPerson := api.persFromMail[email]
	if existingPerson {
		return errors.New("person already exists")
	}
	person := newPerson(email)
	api.People = append(api.People, person)
	api.persFromMail[email] = person
	return nil
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
func (api *Api) AddSkill(role string, title string, desc string,
	parent int32) (uid int32, err error) {

	// Special case when tree is empty
	if api.SkillRoot == -1 {
		uid = api.nextSkill
		api.nextSkill++
		skill := newSkillNode(uid, role, title, desc, -1)
		api.Skills = append(api.Skills, skill)
		api.skillFromId[uid] = skill
		api.SkillRoot = uid
		return
	}
	parentSkill, ok := api.skillFromId[parent]
	if !ok {
		err = errors.New("Unknown parent.")
		return
	}
	if parentSkill.role != CATEGORY {
		err = errors.New("Parent must be a category node")
		return
	}
	uid = api.nextSkill
	api.nextSkill++
	newSkill := newSkillNode(uid, role, title, desc, parentSkill.uid)
	api.skillFromId[uid] = newSkill
	parentSkill.addChild(newSkill.uid)
	return
}

/*
The GivePersonSkill() method adds the given skill into the set of skills the
model holds for that person.  You are only allowed to give people SKILLS, not
CATEGORIES.  An error is generated if either the person or skill given are not
recognized, or you give a person a CATEGORY rather than a SKILL.
*/
func (api *Api) GivePersonSkill(email string, skillId int32) error {
	foundPerson, ok := api.persFromMail[email]
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
	api.SkillHoldings.bind(foundSkill.uid, foundPerson.email)
	return nil
}

func (api *Api) Serialize() (out []byte, err error) {
	return yaml.Marshal(api)
}
