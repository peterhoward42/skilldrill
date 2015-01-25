package model

type uiState struct {
    collapsed []*skillNode
}

func newUiState() *uiState {
    return &uiState{
        collapsed: []*skillNode{},
    }
}

type uiStates struct {
    stateOfPerson map[string]*uiState
}

func newUiStates() *uiStates {
    return &uiStates{
        statesOfPeople: map[string]*uiState{},
    }
}

func (states *uiStates) notifyPersonAdded(emailName string) {
    states.statesOfPeople[emailName] = newUiState()
}
