package model

type modelImpl struct {
	tree     *skillTree
	holdings *holdings
	uiStates *uiStates
}

func newModelImpl() *modelImpl {
	tree := newSkillTree()
	holdings := newHoldings()
	uiStates := newUiStates()
	return &modelImpl{
		tree:     tree,
		holdings: holdings,
		uiStates: uiStates,
	}
}

func (impl *modelImpl) addPerson(emailName string) {
	impl.holdings.notifyPersonAdded(emailName)
	impl.uiStates.notifyPersonAdded(emailName)
}

func (impl *modelImpl) addSkillNode(title string,
	description string, parent int) (skillId int) {
	skillId, skillNode := impl.tree.addSkillNode(title, description, parent)
	impl.holdings.notifySkillAdded(skillNode)
	return
}
