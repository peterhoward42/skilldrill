package model

type treeNode struct {
	title    string
	desc     string
	parent   *treeNode
	children []*treeNode
}

func newTreeNode(title string, desc string, parent *treeNode) *treeNode {
	return &treeNode{
		title:    title,
		desc:     desc,
		parent:   parent,
		children: []*treeNode{},
	}
}

func (parent *treeNode) addChild(child *treeNode) {
    parent.children = append(parent.children, child)
}
