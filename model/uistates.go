package model

import (
	"github.com/peterhoward42/skilldrill/util/sets"
)

/*
The uiState type is a model for the state in which an abstracted
user-experience is in. For example, if some nodes in a tree display of the
skills hiearachy have been collapsed.
*/
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

/*
The uiStates type is a container for a set of uiState objects - corresponding
to a set of people. Each person gets their own state. The type offers a map
between a user's email and their state object.
*/
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
