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
        stateOfPerson: map[string]*uiState{},
    }
}

// Mandated Interface
func (states *uiStates) personAdded(emailName string) {
    states.stateOfPerson[emailName] = newUiState()
}
