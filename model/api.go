/*
The skilldrill model package is a multi-file package that can model a
hierachical set of skills, a set of people who hold some of those skills, and
abstracted user experience states for each person. The api.go file exposes the
Api type, which provides methods for CRUD operations, while the other files
deal with much of the internal workings. The model supports serialization and
de-serialization using
yaml.
*/
package model

import (
	"errors"
	"gopkg.in/yaml.v2"
	"strings"
)

/*
The Api structure is the fundamental type exposed by the skilldrill model
package, and provides CRUD interfaces to do things like adding skills or people
into the model and registering a person as having a particular skill.  All
model editing operations should be done via Api calls rather than accessing the
internal objects directly, so that the integrity of various supplemental look
up tables is preserved.  The design intent is that none of Api fields are
exported, but the reason that some are, is solely to facilitate automated
serialization by yaml.Marshal().
*/
type Api struct {
	SerializeVers int
	Skills        []*skillNode
	People        []*person
	SkillRoot     int            // root of taxonomy tree (skill.Uid)
	SkillHoldings *skillHoldings // who has what skill?
	NextSkill     int
	UiStates      map[string]*uiState
	// Supplemental, (duplicate) data for quick lookups
	skillFromId  map[int]*skillNode
	persFromMail map[string]*person
}

// The function NewApi() is a (compulsory) constructor for an initialized, but
// empty Api struct.
func NewApi() *Api {
	return &Api{
		SerializeVers: 1,
		Skills:        make([]*skillNode, 0),
		People:        make([]*person, 0),
		SkillRoot:     -1,
		SkillHoldings: newSkillHoldings(),
		NextSkill:     1,
		UiStates:      make(map[string]*uiState),
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

//--------------------------------------------------------------------------
// Methods For Adding things to the model
//--------------------------------------------------------------------------

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
	// be sure to keep this symmetrical with RemovePerson()
	incomer := newPerson(email)
	api.People = append(api.People, incomer)
	api.persFromMail[email] = incomer
	api.SkillHoldings.registerPerson(email)
	api.UiStates[email] = newUiState()
	return nil
}

/*
The AddSkill() method adds a skill into the model's hierachy of skills.  You
specify the skill in terms of description and title strings. These strings
should describe how they additionally qualify their context in the hierachy,
and should not duplicate context information.  You specify the tree location by
providing the Uid of the parent skill, and the new Uid for the added skill is
returned.  The role parameter should be one of the constants Skill or Category.
When the skill tree is empty, this skill will be added as the root, and the
parentUid parameter is ignored.  Errors are generated if you attempt to add a
skill to a node that is not a Category, or if the parent skill you provide is
not recognized.
*/
func (api *Api) AddSkill(role string, title string, desc string,
	parent int) (uid int, err error) {

	// Be sure to keep this symmetrical with RemoveSkill

	// Sanitize parent except when adding root skill
	if api.SkillRoot != -1 {
		parentSkill, ok := api.skillFromId[parent]
		if !ok {
			err = errors.New(UnknownParent)
			return
		}
		if parentSkill.Role != Category {
			err = errors.New(ParentNotCategory)
			return
		}
	}
	uid = api.NextSkill
	api.NextSkill++
	newSkill := newSkillNode(uid, role, title, desc, parent, api)
	// Note we keep the children - in alphabetical order of title
	api.Skills = append(api.Skills, newSkill)
	api.skillFromId[uid] = newSkill
	api.SkillHoldings.registerSkill(uid)

	if api.SkillRoot == -1 {
		api.SkillRoot = uid
		return
	}
	parentSkill := api.skillFromId[parent]
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
func (api *Api) GivePersonSkill(email string, skillId int) (err error) {
	if err = api.tweakParams(&email, &skillId); err != nil {
		return
	}
	foundSkill := api.skillFromId[skillId]
	if foundSkill.Role == Category {
		err = errors.New(CannotBestowCategory)
		return
	}
	foundPerson := api.persFromMail[email]
	api.SkillHoldings.bind(foundSkill.Uid, foundPerson.Email)
	return
}

//--------------------------------------------------------------------------
// Methods For Editing the UXP State
//--------------------------------------------------------------------------

/*
The CollapseSkill() method operates on the part of the model that represents
the abstracted user experience. In this case to collapse a node in the tree
display of skills hierachy. Errors are generated when either the person or the
skill is not recognized.
*/
func (api *Api) CollapseSkill(email string, skillId int) (err error) {
	if err = api.tweakParams(&email, &skillId); err != nil {
		return
	}
	foundSkill := api.skillFromId[skillId]
	api.UiStates[email].collapseNode(foundSkill)
	return
}

//--------------------------------------------------------------------------
// Getter Style Methods
//--------------------------------------------------------------------------
/*
The method SkillWording() returns the title and description of the given skill.
The description is provided in three different forms: The description in
isolation of the skill node, this same description tacked on to the end of a
description of the ancestry, and this ancestry part isolated. Can generate the
UnknownSkill error.
*/
func (api *Api) SkillWording(skillId int) (title string, desc string,
	descInContext string, contextAlone string, err error) {
	if err = api.tweakParams(nil, &skillId); err != nil {
		return
	}
	foundSkill := api.skillFromId[skillId]
	treeOps := &skillTreeOps{api}
	title, desc, descInContext, contextAlone = treeOps.skillWording(foundSkill)
	return
}

/*
The method PeopleWithSkill() provides a list of the people (email address) who
hold the given skill. Can generate the following errors: UnknownSkill,
CannotBestowCategory.
*/
func (api *Api) PeopleWithSkill(skillId int) (emails []string, err error) {
	if err = api.tweakParams(nil, &skillId); err != nil {
		return
	}
	if api.skillFromId[skillId].Role == Category {
		err = errors.New(CannotBestowCategory)
		return
	}
	emails = api.SkillHoldings.PeopleWithSkill[skillId].AsSlice()
	return
}

/*
The method PersonExists() returns true if the given person is registered.
*/
func (api *Api) PersonExists(email string) bool {
	_, exists := api.persFromMail[email]
	return exists
}

/*
The method PersonHasSkill() returns true if the given person is registered as
having the given skill.  Can generate the following errors: UnknownSkill,
UnknownPerson, CannotBestowCategory.
*/
func (api *Api) PersonHasSkill(email string, skillId int) (
	hasSkill bool, err error) {
	if err = api.tweakParams(&email, &skillId); err != nil {
		return
	}
	if api.skillFromId[skillId].Role == Category {
		err = errors.New(CannotBestowCategory)
		return
	}
	sh := api.SkillHoldings
	sop := sh.SkillsOfPerson
	set := sop[email]
	present := set.Contains(skillId)
	_ = present
	hasSkill = api.SkillHoldings.SkillsOfPerson[email].Contains(skillId)
	return
}

/*
The method EnumerateTree() provides a list of skill Uids in the order they
should appear when displaying the tree. It is person-specific, and omits the
nodes that have been collapsed (using CollapseSkill()) - including their
children. Can generate the UnknownPerson error.
*/
func (api *Api) EnumerateTree(email string) (skills []int,
	depths []int, err error) {
	if err = api.tweakParams(&email, nil); err != nil {
		return
	}
	treeOps := &skillTreeOps{api}
	collapsedNodes := api.UiStates[email].CollapsedNodes
	skills, depths = treeOps.enumerateTree(collapsedNodes)
	return
}

//--------------------------------------------------------------------------
// Methods That Change Existing Content
//--------------------------------------------------------------------------

/*
The method ReParentSkill() moves a skill node and all its children to a
different position in the tree. The new parent given must be a skill node with
the CATEGORY role. The following errors can be generated: UnknownSkill,
IllegalWithRoot, and ParentNotCategory.
*/
func (api *Api) ReParentSkill(toMove int, newParent int) (err error) {
	if err = api.tweakParams(nil, &toMove); err != nil {
		return
	}
	if err = api.tweakParams(nil, &newParent); err != nil {
		return
	}
	if toMove == api.SkillRoot {
		return errors.New(IllegalWithRoot)
	}
	childSkill := api.skillFromId[toMove]
	oldParentSkill := api.skillFromId[childSkill.Parent]
	newParentSkill := api.skillFromId[newParent]

	if newParentSkill.Role != Category {
		return errors.New(ParentNotCategory)
	}
	oldParentSkill.removeChild(toMove)
	newParentSkill.addChild(toMove)
	childSkill.Parent = newParent
	return
}

/*
The RemovePerson() method removes a previously registered person from the model
in terms of the user name part of their email address. It is an error to pass
in a person that does not exist in the model. The email address is coerced to
lowercase.
*/
func (api *Api) RemovePerson(email string) (err error) {
	if err = api.tweakParams(&email, nil); err != nil {
		return
	}
	// be sure to keep this symmetrical with AddPerson()
	departingPerson := api.persFromMail[email]
	oldList := api.People
	api.People = []*person{}
	for _, incumbentPerson := range oldList {
		if incumbentPerson != departingPerson {
			api.People = append(api.People, incumbentPerson)
		}
	}
	delete(api.persFromMail, email)
	api.SkillHoldings.UnRegisterPerson(*departingPerson)
	delete(api.UiStates, email)
	return
}

/*
The RemoveSkill() method removes a skill from the model's hierachy of skills.
It can generate the following errors: UnknownSkill,
CannotRemoveSkillWithChildren, CannotRemoveRootSkill.
*/
func (api *Api) RemoveSkill(skillId int) (err error) {
	// Be sure to keep this symmetrical with RemoveSkill
	if err = api.tweakParams(nil, &skillId); err != nil {
		return
	}
	// The order of the following 2 tests makes it easier to design tests.
	if skillId == api.SkillRoot {
		err = errors.New(CannotRemoveRootSkill)
		return
	}
	departingSkill := api.skillFromId[skillId]
	if len(departingSkill.Children) != 0 {
		err = errors.New(CannotRemoveSkillWithChildren)
		return
	}
	parentSkill := api.skillFromId[departingSkill.Parent]
	parentSkill.removeChild(skillId)
	oldList := api.Skills
	api.Skills = []*skillNode{}
	for _, incumbentSkill := range oldList {
		if incumbentSkill != departingSkill {
			api.Skills = append(api.Skills, incumbentSkill)
		}
	}
	delete(api.skillFromId, skillId)
	api.SkillHoldings.UnRegisterSkill(*departingSkill)
	return
}

//--------------------------------------------------------------------------
// Serialize Methods
//--------------------------------------------------------------------------

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

//--------------------------------------------------------------------------
// Module Private Methods
//--------------------------------------------------------------------------

/*
The method tweakParams(), receives either or both of an email and a skill Uid,
and coerces the email when given into lowercase, and then ensures the email is
one known to the model, and the skillUid is legitimate. It can return either
of the errors: UnknownPerson or UnknownSkill.
*/
func (api *Api) tweakParams(email *string, skillId *int) (err error) {
	if email != nil {
		// coerce caller's email to lower case
		*email = strings.ToLower(*email)
		_, ok := api.persFromMail[*email]
		if !ok {
			return errors.New(UnknownPerson)
		}
	}
	if skillId != nil {
		_, ok := api.skillFromId[*skillId]
		if !ok {
			return errors.New(UnknownSkill)
		}
	}
	return
}

// The method titleFromId() exists to satisfy the titleMapper interface.
func (api *Api) titleFromId(skillUid int) (title string) {
	return api.skillFromId[skillUid].Title
}

/*
The SetSkillTitle() method replaces the given skill's title with the text
given. Can generate the following errors: SkillUnknown error, TooLong.
*/
func (api *Api) SetSkillTitle(skillId int, newTitle string) (err error) {
	if err = api.tweakParams(nil, &skillId); err != nil {
		return
	}
	skill := api.skillFromId[skillId]
	if len(newTitle) > MaxSkillTitle {
		err = errors.New(TooLong)
		return
	}
	skill.Title = newTitle
	return
}

/*
The SetSkillDesc() method replaces the given skill's description with the text
given. Can generate the following errors: SkillUnknown error, TooLong.
*/
func (api *Api) SetSkillDesc(skillId int, newDesc string) (err error) {
	if err = api.tweakParams(nil, &skillId); err != nil {
		return
	}
	if len(newDesc) > MaxSkillDesc {
		err = errors.New(TooLong)
		return
	}
	skill := api.skillFromId[skillId]
	skill.Desc = newDesc
	return
}
