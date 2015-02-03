/*
The model package is a data model and CRUD interface for data that represents a
hierachical set of skills, and a set of people who hold some of these skills.
You can for example add or remove skills, add or remove people and register
people as having a one of the skills. The external interface is exposed by the
methods belonging to the Api type.
*/
package model

import (
	"errors"
)

/*
The Api type exposes the exernal API to the model package. In most cases, the
api methods are simple wrappers to sister methods in the model type which take
care of validating the input parameters like email names and skill Uids. This
frees all the other modules from checking method parameters, and thus keeps
them simpler.
*/
type Api struct {
	model *model
}

func NewApi() *Api {
	return &Api{
		model: newModel(),
	}
}

//----------------------------------------------------------------------------
// Add methods
//----------------------------------------------------------------------------

/*
The AddPerson method adds a person to the model. A person is defined throughout
the model by the name part of their email address. Errors: PersonExists.
*/
func (api *Api) AddPerson(emailName string) (err error) {
	if api.model.holdings.personExists(emailName) {
		err = errors.New(PersonExists)
		return
	}
	api.model.addPerson(emailName)
	return
}

/*
The AddSkillNode method adds a skill to the model, as a child of the given
parent skill. The method returns a unique ID for the skill (Uid) by which it
may be referred to subsequently. If the model is empty, the new node is adopted
as the root for the tree, and the parent parameter is ignored. The role of
skill nodes in the tree as branches or leaves remains open until a person
registers as having a skill. At that point, the skill becomes fixed as a leaf,
and children may not be added to it subsequently. Errors: UnknownSkill
(parent), IllegalForLeaf.
*/
func (api *Api) AddSkillNode(title string, description string,
	parent int) (skillId int, err error) {
	if api.model.tree.treeIsEmpty() {
		skillId = api.model.addRootSkillNode(title, description)
		return
	}
	if api.model.tree.skillExists(parent) == false {
		err = errors.New(UnknownSkill)
		return
	}
	parentNode := api.model.tree.nodeFromUid[parent]
	if api.model.holdings.someoneHasThisSkill(parentNode) {
		err = errors.New(IllegalForHeldSkill)
		return
	}
	skillId = api.model.addChildSkillNode(title, description, parent)
	return
}

/*
The GivePersonSkill method registers the given person as having the given
skill. You cannot assign skills in the tree that have children to people.
Errors: UnknownSkill, UnknownPerson, DisallowedForAParent.
*/
func (api *Api) GivePersonSkill(emailName string, skillId int) (err error) {
	if api.model.holdings.personExists(emailName) == false {
		err = errors.New(UnknownPerson)
		return
	}
	if api.model.tree.skillExists(skillId) == false {
		err = errors.New(UnknownSkill)
		return
	}
	skill := api.model.tree.nodeFromUid[skillId]
	api.model.givePersonSkill(skill, emailName)
	return
}

//----------------------------------------------------------------------------
// UiState editing (in model space)
//----------------------------------------------------------------------------

/*
The ToggleSkillCollapsed method marks the given skill in the tree as being
collapsed for the given person. (or vice versa depending on the current state).
It is illegal to call it on a node with no children. Errors: UnknownSkill,
UnknownPerson, IllegalWhenNoChildren
*/
func (api *Api) ToggleSkillCollapsed(
	emailName string, skillId int) (err error) {
	if api.model.holdings.personExists(emailName) == false {
		err = errors.New(UnknownPerson)
		return
	}
	if api.model.tree.skillExists(skillId) == false {
		err = errors.New(UnknownSkill)
		return
	}
	skill := api.model.tree.nodeFromUid[skillId]
	if skill.hasChildren() == false {
		err = errors.New(IllegalWhenNoChildren)
		return
	}
	api.model.toggleSkillCollapsed(emailName, skill)
	return
}

//----------------------------------------------------------------------------
// Queries
//----------------------------------------------------------------------------

func (api *Api) PersonExists(emailName string) bool {
	return api.model.personExists(emailName)
}

func (api *Api) SkillExists(skillId int) bool {
	return api.model.skillExists(skillId)
}

func (api *Api) TitleOfSkill(skillId int) (title string, err error) {
	if api.model.tree.skillExists(skillId) == false {
		err = errors.New(UnknownSkill)
		return
	}
	title = api.model.titleOfSkill(skillId)
	return
}

func (api *Api) PersonHasSkill(skillId int, email string) (
    hasSkilll bool, err error) {
	if api.model.skillExists(skillId) == false {
		err = errors.New(UnknownSkill)
		return
	}
	if api.model.personExists(emailName) == false {
		err = errors.New(UnknownPerson)
		return
	}
    return api.model.personHasSkill(skillId, email)
}

/*
The EnumerateTree method provides a linear sequence of TreeDisplayItem which
can be used to used to render the skill tree. It is personalised to a given
emailName, and will have omitted the children of any skill nodes the person has
collapsed.  Errors: UnknownPerson
func (api *Api) EnumerateTree(emailName string) (
	displayRows []TreeDisplayItem, err error) {
	if api.model.holdings.personExists(emailName) == false {
		err = errors.New(UnknownPerson)
		return
	}
	displayRows = api.model.enumerateTree(emailName)
	return
}
*/
