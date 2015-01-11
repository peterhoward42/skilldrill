/*
The skilldrill model package is a multi-file package that can model a hierachical
set of skills and a set of people who hold some of those skills. The api.go file
exposes the Api type, which provides methods for CRUD operations, while the other
files deal with much of the internal workings. The model supports serialization
and de-serialization using a yaml form.
*/
package model

import (
	"errors"
	"gopkg.in/yaml.v2"
	"strings"
)

// The Api structure is the fundamental type exposed by the skilldrill model
// package, and provides CRUD interfaces to do things like adding skills or
// people into the model and registering a person as having a particular skill.
// All model editing operations should be done via Api calls rather than
// accessing the internal objects directly, so that the integrity of various
// supplemental look up tables is preserved.
// The design intent is that none of Api fields are exported, but the reason
// that some are, is solely to facilitate automated serialization by
// yaml.Marshal().
type Api struct {
	SerializeVers  int
	Skills         []*skillNode
	People         []*person
	SkillRoot      int            // root of taxonomy tree (skill.Uid)
	SkillHoldings  *skillHoldings // who has what skill?
	NextSkill      int
	UiStates       map[string]*uiState
	// Supplemental, (duplicate) data for quick lookups
	skillFromId  map[int]*skillNode
	persFromMail map[string]*person
}

// The function NewApi() is a (compulsory) constructor for an initialized, but
// empty Api struct.
func NewApi() *Api {
	return &Api{
		SerializeVers:  1,
		Skills:         make([]*skillNode, 0),
		People:         make([]*person, 0),
		SkillRoot:      -1,
		SkillHoldings:  newSkillHoldings(),
		NextSkill:      1,
		UiStates:       make(map[string]*uiState),
		// Supplemental fields
		skillFromId:  make(map[int]*skillNode),
		persFromMail: make(map[string]*person),
	}
}

// The function NewFromSerialized() is a factory for an Api based on
// content previously serialized using the Api.Serialize() method.
func NewFromSerialized(in []byte) (api *Api, err error) {
	api = NewApi()
	err = yaml.Unmarshal(in, api)
	if err != nil {
		return
	}
	api.finishBuildFromDeSerialize()
	return
}

// The AddPerson() method adds a person to the model in terms of the user name
// part of their email address. It is an error to add a person that already
// exists in the model. The email address is coerced to lowercase.
func (api *Api) AddPerson(email string) (err error) {
	// disallow duplicate additions
	email = strings.ToLower(email)
	_, ok := api.persFromMail[email]
	if ok {
		return errors.New(PersonExists)
	}
	incomer := newPerson(email)
	api.People = append(api.People, incomer)
	api.persFromMail[email] = incomer
    api.UiStates[email] = newUiState()
	return nil
}

/*
The AddSkill() method adds a skill into the model's hierachy of skills.  You
specify the skill in terms of description and title strings. These strings should
describe how they additionally qualify their context in the hierachy, and should
not duplicate context information.  You specify the tree location by providing
the Uid of the parent skill, and the new Uid for the added skill is returned.
The role parameter should be one of the constants Skill or Category.  When the
skill tree is empty, this skill will be added as the root, and the parentUid
parameter is ignored.  Errors are generated if you attempt to add a skill to a
node that is not a Category, or if the parent skill you provide is not
recognized.
*/
func (api *Api) AddSkill(role string, title string, desc string,
	parent int) (uid int, err error) {

	// Special case when tree is empty
	if api.SkillRoot == -1 {
		uid = api.NextSkill
		api.NextSkill++
		skill := newSkillNode(uid, role, title, desc, -1)
		api.Skills = append(api.Skills, skill)
		api.skillFromId[uid] = skill
		api.SkillRoot = uid
		return
	}
	parentSkill, ok := api.skillFromId[parent]
	if !ok {
		err = errors.New(UnknownParent)
		return
	}
	if parentSkill.Role != Category {
		err = errors.New(ParentNotCategory)
		return
	}
	uid = api.NextSkill
	api.NextSkill++
	newSkill := newSkillNode(uid, role, title, desc, parentSkill.Uid)
	api.Skills = append(api.Skills, newSkill)
	api.skillFromId[uid] = newSkill
	parentSkill.addChild(newSkill.Uid)
	return
}

/*
The GivePersonSkill() method adds the given skill into the set of skills the
model holds for that person.  You are only allowed to give people Skill, not
CATEGORIES.  An error is generated if either the person or skill given are not
recognized, or you give a person a Category rather than a Skill. The email you
provide is lower-cased before it is used.
*/
func (api *Api) GivePersonSkill(email string, skillId int) error {
	email = strings.ToLower(email)
	foundPerson, ok := api.persFromMail[email]
	if !ok {
		return errors.New(UnknownPerson)
	}
	foundSkill, ok := api.skillFromId[skillId]
	if !ok {
		return errors.New(UnknownSkill)
	}
	if foundSkill.Role == Category {
		return errors.New(CategoryDisallowed)
	}
	api.SkillHoldings.bind(foundSkill.Uid, foundPerson.Email)
	return nil
}

/*
The CollapseSkill() method operates on the part of the model that represents the
abstracted user experience. In this case to collapse a node in the tree display
of skills hierachy. Errors are generated when either the person or the skill is
not recognized.
*/
func (api *Api) CollapseSkill(email string, skillId int) error {
	email = strings.ToLower(email)
	_, ok := api.persFromMail[email]
	if !ok {
		return errors.New(UnknownPerson)
	}
	foundSkill, ok := api.skillFromId[skillId]
	if !ok {
		return errors.New(UnknownSkill)
	}
	api.UiStates[email].collapseNode(foundSkill)
	return nil
}

/*
The function Serialize() makes a machine-readable representation of the Api
object and packages it into a slice of bytes. See also NewFromSerialized().
*/
func (api *Api) Serialize() (out []byte, err error) {
	return yaml.Marshal(api)
}

/*
The function finishBuildFromDeSerialize() takes the state of an Api object that
has been partly initialized from de-serialization, and builds the supplemental
fields required. These are mainly look up tables for convenience and speed.
*/
func (api *Api) finishBuildFromDeSerialize() {
	for _, skill := range api.Skills {
		uid := skill.Uid
		api.skillFromId[uid] = skill
	}
	for _, person := range api.People {
		email := person.Email
		api.persFromMail[email] = person
	}
}
