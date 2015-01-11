package model

// This enumerated type provides a classification for the mutually exclusive
// roles that a skillNode may take.
const (
	Skill    = "SKL"
	Category = "CAT"
)

// These constants provide a set of human-readable error message strings, with
// machine-readable names.
const (
	PersonExists       = "Person exists."
	UnknownParent      = "Unknown parent."
	ParentNotCategory  = "Parent must be a category node."
	UnknownPerson      = "Person does not exist."
	UnknownSkill       = "Skill does not exist."
	CategoryDisallowed = "Cannot give someone a CATEGORY skill."
)
