package model

const (
	SKILL = iota
	CATEGORY
)

type skillNode struct {
	uid      int64
	role     int // SKILL | CATEGORY
	title    string
	desc     string
	parent   *skillNode
	children []*skillNode
}

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

func (parent *skillNode) addChild(child *skillNode) {
	parent.children = append(parent.children, child)
}
