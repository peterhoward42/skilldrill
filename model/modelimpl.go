package model

type modelImpl struct {
    tree    *skillTree
    holdings *holdings
    uistates *uiStates
}

func newModelImpl() *modelImpl {
    return &modelImpl{
        tree:    newSkillTree(),
        holdings: newHoldings(),
        uistates: newUiStates(),
    }
}
