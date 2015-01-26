package model

// These constants provide a set of human-readable error message strings, with
// machine-readable names.
const (
	DisallowedForAParent  = "Illegal for node with children."
	CannotRemoveSkillHeld = "Cannot remove a skill that people have."
	IllegalForLeaf        = "Cannot add child to a <held> skill."
	IllegalWithRoot       = "Cannot be done with root skill."
	PersonExists          = "Person exists."
	UnknownPerson         = "Person does not exist."
	UnknownSkill          = "Skill does not exist."
)
