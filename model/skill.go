package model

import (
	"sort"
)

type skillNode struct {
	title    string
	desc     string
	uid      int
	parent   *skillNode
	children []*skillNode
}

func newSkillNode(title string, desc string, uid int,
	parent *skillNode) (createdNode *skillNode) {
	return &skillNode{
		title:  title,
		desc:   desc,
		uid:    uid,
		parent: parent,
	}
}

func (parent *skillNode) addChild(newChild *skillNode) {
	// This maintains children in alphabetical order
	titles := []string{}
	children := parent.children
	children = append(children, newChild)
	title2Node := map[string]*skillNode{}
	for _, child := range children {
		titles = append(titles, child.title)
		title2Node[child.title] = child
	}
	sort.Strings(titles)
	// make new slice - in sorted order
	parent.children = []*skillNode{}
	for _, title := range titles {
		parent.children = append(parent.children, title2Node[title])
	}
}
