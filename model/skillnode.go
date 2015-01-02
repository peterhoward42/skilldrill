package model

type skillNode struct {
	uid      int64
	title    string
	desc     string
	parent   *skillNode
	children []*skillNode
}

func newSkillNode(uid int64, title string, desc string,
	parent *skillNode) *skillNode {
	return &skillNode{
		uid:      uid,
		title:    title,
		desc:     desc,
		parent:   parent,
		children: []*skillNode{},
	}
}

func (parent *skillNode) addChild(child *skillNode) {
	parent.children = append(parent.children, child)
}
