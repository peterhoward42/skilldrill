package model

import (
	"github.com/peterhoward42/skilldrill/util/sets"
)

type uiState struct {
	collapsed *sets.SetOfInt
}

func newUiState() *uiState {
	return &uiState{
		collapsed: sets.NewSetOfInt(),
	}
}

func (state *uiState) toggleCollapsed(skill *skillNode) {
	state.collapsed.TogglePresenceOf(skill.uid)
}

type uiStates struct {
	stateOfPerson map[string]*uiState
}

func newUiStates() *uiStates {
	return &uiStates{
		stateOfPerson: map[string]*uiState{},
	}
}

func (states *uiStates) notifyPersonAdded(emailName string) {
	states.stateOfPerson[emailName] = newUiState()
}
