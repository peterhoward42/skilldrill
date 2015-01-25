package model

type modelImpl struct {
    tree    *skillTree
    holdings *holdings
    uiStates *uiStates
    observers []changeObserver
}

func newModelImpl() *modelImpl {
    tree := newSkillTree()
    holdings := newHoldings()
    uiStates := newUiStates()
    return &modelImpl{
        tree:    tree,
        holdings: holdings,
        uiStates: uiStates,
        observers: []changeObserver{tree, holdings, uiStates},
    }
}

func (impl *modelImpl) addPerson(emailName string) {
    impl.propagatePersonAdded(emailName)
}

func (impl *modelImpl) addSkillNode(title string, 
    description string, parent int) (uid int) {
    skillNode := impl.tree.addSkillNode(title, description, parent)
    impl.propagateSkillNodeAdded(skillNode)
}

func (impl *modelImpl) propagatePersonAdded(emailName string) {
    for _, observer := range impl.observers {
        observer.personAdded(emailName)
    }
}
