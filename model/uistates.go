package model

type uiState struct {
    collapsed []*skillNode
}

type uiStates struct {
    stateOfPerson map[string]uiState
}

func newUiStates() *uiStates {
    return &uiStates{
        stateOfPerson: make(map[string]uiState),
    }
}
