package model

type model struct {
	tree     *tree
	holdings *holdings
	uiStates *uiStates
}

func newModel() *modelImpl {
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
}

func (model *model) addChildSkillNode(title string,
	description string, parent int) (skillId int) {
    skillNode, skillId := model.tree.addChildSkillNode(title, description)
    model.holdings.notifySkillAdded(skillNode)
}

//---------------------------------------------------------------------------
// Query operations
//---------------------------------------------------------------------------

//---------------------------------------------------------------------------
// Edit operations
//---------------------------------------------------------------------------
