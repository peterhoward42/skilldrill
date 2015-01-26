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

func (model *model) addSkillNode(title string,
	description string, parent int) (skillId int) {
	skillId, skillNode := model.tree.addSkillNode(title, description, parent)
	model.holdings.notifySkillAdded(skillNode)
	return
}

//---------------------------------------------------------------------------
// Query operations
//---------------------------------------------------------------------------

func (model *model) personExists(emailName string) bool {
    return model.holdings.personExists(emailName)
}

func (model *model) treeIsEmpty() bool {
    return model.tree.isEmpty()
}

//---------------------------------------------------------------------------
// Edit operations
//---------------------------------------------------------------------------
