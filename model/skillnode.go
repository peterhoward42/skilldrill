package model

import ()

// This enumerated type provides a classification for the mutually exclusive
// role that a skillNode may take.
const (
	SKILL    = "SKL"
	CATEGORY = "CAT"
)

/*
The skillNode type models a node in the skill hierachy. Only the leaf-level nodes
are permitted to be SKILLS. The others must have the role CATEGORY. The
descriptive strings should provide only a qualification for their specialism with
respect to their parent category, and should not duplicate this information.
The design intent is that none of Api fields are exported, but the reason
that some are, is solely to facilitate automated serialization by
yaml.Marshal().
*/
type skillNode struct {
	Uid      int
	Role     string // SKILL | CATEGORY
	Title    string
	Desc     string
	Parent   int
	Children []int
}

// Compulsory constructor.
func newSkillNode(uid int, role string, title string, desc string,
	parent int) *skillNode {
	return &skillNode{
		Uid:      uid,
		Role:     role,
		Title:    title,
		Desc:     desc,
		Parent:   parent,
		Children: []int{},
	}
}

// The method addChild() adds the given skillNode into the tree as a child of the
// given parent.
func (parent *skillNode) addChild(child int) {
	parent.Children = append(parent.Children, child)
}
