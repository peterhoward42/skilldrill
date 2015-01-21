package model

// This enumerated type provides a classification for the mutually exclusive
// roles that a skillNode may take.
const (
	Skill    = "SKL"
	Category = "CAT"
)

// These constants specify the maximum length allowed for various fields.
const (
	MaxSkillTitle int = 30
	MaxSkillDesc  int = 400
)

// These constants provide a set of human-readable error message strings, with
// machine-readable names.
const (
	CannotBestowCategory          = "Cannot give someone a CATEGORY skill."
	CannotRemoveRootSkill         = "Cannot remove the root skill."
	CannotRemoveSkillWithChildren = "Cannot remove skill with children"
	IllegalWithRoot               = "Cannot be done with root skill."
	ParentNotCategory             = "Parent must be a category node."
	PersonExists                  = "Person exists."
	TooLong                       = "String is too long."
	UnknownParent                 = "Unknown parent."
	UnknownPerson                 = "Person does not exist."
	UnknownSkill                  = "Skill does not exist."
)
