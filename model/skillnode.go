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
*/
type skillNode struct {
	uid      int64
	role     string // SKILL | CATEGORY
	title    string
	desc     string
	parent   *skillNode
	children []*skillNode
}

// Compulsory constructor.
func newSkillNode(uid int64, role string, title string, desc string,
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

// The method childUids() provides a slice of UIDs of this skill's children.
func (s *skillNode) childUids() (ids []int64) {
	for _, child := range s.children {
		ids = append(ids, child.uid)
	}
	return
}
