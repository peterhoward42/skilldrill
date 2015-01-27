package model

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

/*
The EnumerateTree method provides a linear sequence of TreeDisplayItem which
can be used to used to render the skill tree. It is personalised to a given
emailName, and will have omitted the children of any skill nodes the person has
collapsed.

I have created this intermediate, pass-through wrapper in case it is needed to
inject additional data sources downstream.
*/
func (model *model) EnumerateTree(emailName string) (
    displayRows []TreeDisplayItem) {
    collapsed := model.uiStates[emailName].collapsed
    return = api.model.tree.enumerateTree(collapsed)
}

//---------------------------------------------------------------------------
// UiState operations (in model space)
//---------------------------------------------------------------------------

func (model *model) toggleSkillCollapsed(emailName string, skill *skillNode) {
	model.uiStates.stateOfPerson[emailName].toggleCollapsed(skill)
}
