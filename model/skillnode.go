package model

import ()

// This enumerated type provides a classification for the mutually exclusive
// role that a skillNode may take.
const (
	SKILL = iota
	CATEGORY
)

/*
The skillNode type models a node in the skill hierachy. Only the leaf-level nodes
are permitted to be SKILLS. The others must have the role CATEGORY. The
descriptive strings should provide only a qualification for their specialism with
respect to their parent category, and should not duplicate this information.
*/
type skillNode struct {
	uid      int64
	role     int // SKILL | CATEGORY
	title    string
	desc     string
	parent   *skillNode
	children []*skillNode
}

// Compulsory constructor.
func newSkillNode(uid int64, role int, title string, desc string,
	parent *skillNode) *skillNode {
	return &skillNode{
		uid:      uid,
		role:     role,
		title:    title,
		desc:     desc,
		parent:   parent,
		children: []*skillNode{},
	}
}

// The method addChild() adds the given skillNode into the tree as a child of the
// given parent.
func (parent *skillNode) addChild(child *skillNode) {
	parent.children = append(parent.children, child)
}
