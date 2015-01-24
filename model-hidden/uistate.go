package model

import (
	"github.com/peterhoward42/skilldrill/util/sets"
)

/*
The uiState() type is the model that represents a state that the abstracted
user experience can be in. For example, which of the nodes in the skills tree
are collapsed.  The design intent is that none of Api fields are exported, but
the reason that some are, is solely to facilitate automated serialization by
yaml.Marshal().
*/
type uiState struct {
	CollapsedNodes *sets.SetOfInt
}

// Compulsory constructor.
func newUiState() *uiState {
	return &uiState{CollapsedNodes: sets.NewSetOfInt()}
}

/*
The function collapseNode() adds the given skillNode's Uid to the set of
collapsed nodes held by the uiState object. We use setters and getters for the
collapse and expand operations so that we can control side effects elsewhere in
the tree if required. We use a *skillNode parameter rather than the skill Uid
directly, so as to push responsibility for validating the skillNode externally.
*/
func (s *uiState) collapseNode(node *skillNode) {
	s.CollapsedNodes.Add(node.Uid)
}

func (s *uiState) NotifySkillIsRemoved(skillId int) {
	s.CollapsedNodes.RemoveIfPresent(skillId)
}
