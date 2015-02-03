package model

/*
The model type is the tip of the modelling pyramid and contains a set of
subsidiary models - like one to model the skill hiearchy, another to model who
holds which skill etc. The type provides methods for CRUD operations like
adding a person or allocating a skill to a person. The model implements its
methods for the most part by delegating smaller operations to the subsidiary
models.  It is the model that is responsible for propogating changes between
the subsidiary models, so that the subsidiary models in turn can have minimal
scope and coupling. The model methods do NOT check the legitimacy of the
parameters provided and will panic if they are wrong. For example if an email
address provided is one that is known to the system.
*/
type model struct {
	tree     *tree
	holdings *holdings
	uiStates *uiStates
}

func newModel() *model {
	tree := newTree()
	holdings := newHoldings()
	uiStates := newUiStates()
	return &model{
		tree:     tree,
		holdings: holdings,
		uiStates: uiStates,
	}
}

//---------------------------------------------------------------------------
// Add operations
//---------------------------------------------------------------------------

func (model *model) addPerson(emailName string) {
	model.holdings.notifyPersonAdded(emailName)
	model.uiStates.notifyPersonAdded(emailName)
}

func (model *model) addRootSkillNode(title string,
	description string) (skillId int) {
	skillNode, skillId := model.tree.addRootSkillNode(title, description)
	model.holdings.notifySkillAdded(skillNode)
	return
}

func (model *model) addChildSkillNode(title string,
	description string, parent int) (skillId int) {
	skillNode, skillId := model.tree.addChildSkillNode(title,
		description, parent)
	model.holdings.notifySkillAdded(skillNode)
	return
}

func (model *model) givePersonSkill(skill *skillNode, emailName string) {
	model.holdings.givePersonSkill(skill, emailName)
}

//---------------------------------------------------------------------------
// Query operations
//---------------------------------------------------------------------------

func (model *model) treeIsEmpty() bool {
	return model.tree.treeIsEmpty()
}
func (model *model) personExists(emailName string) bool {
	return model.holdings.personExists(emailName)
}

func (model *model) skillExists(skillId int) bool {
	return model.holdings.skillExists(model.tree.nodeFromUid[skillId])
}


func (model *model) skillNode(skillId int) (skillNode *skillNode) {
	return model.tree.nodeFromUid[skillId]
}


func (model *model) personHasSkill(skillId int, email string) (
    hasSkilll bool) {
    return model.holdings.personHasSkill(
        model.tree.nodeFromUid[skillId], email)
}
    
    
func (model *model) someoneHasThisSkill(skillNode *skillNode) bool {
	return model.holdings.someoneHasThisSkill(skillNode)
}

func (model *model) titleOfSkill(skillId int) (title string) {
	return model.tree.titleOfSkill(skillId)
}

/*
The EnumerateTree method provides a linear sequence of TreeDisplayItem which
can be used to used to render the skill tree. It is personalised to a given
emailName, and will have omitted the children of any skill nodes the person has
collapsed.

I have created this intermediate, pass-through wrapper in case it is needed to
inject additional data sources downstream.
func (model *model) EnumerateTree(emailName string) (
	displayRows []TreeDisplayItem) {
	collapsed := model.uiStates[emailName].collapsed
	return api.model.tree.enumerateTree(collapsed)
}
*/

//---------------------------------------------------------------------------
// UiState operations (in model space)
//---------------------------------------------------------------------------

func (model *model) toggleSkillCollapsed(emailName string, skill *skillNode) {
	model.uiStates.stateOfPerson[emailName].toggleCollapsed(skill)
}
