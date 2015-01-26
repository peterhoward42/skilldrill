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
The Api type exposes the exernal API to the model package. When pre existing
skill Uids or email addresses are used as parameters to the Api methods, these
are validated at the Api level, so that the other modules can be simpler and
clearer.
*/
type Api struct {
	model *model
}

func NewApi() *Api {
	return &Api{
		model: newModel(),
	}
}

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
	if api.model.holdings.skillExists(parent) == false {
		err = errors.New(UnknownSkill)
        return
	}
	if api.model.tree.skillIsALeaf(parent) {
		err = errors.New(IllegalForLeaf)
        return
	}
	skillId = api.model.addChildSkillNode(title, description, parent)
	return
}
