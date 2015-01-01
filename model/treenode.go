package model

type TreeNode struct {
	title    string
	desc     string
	parent   *TreeNode
	children []*TreeNode
}

func NewTreeNode(title string, desc string, parent *TreeNode) *TreeNode {
	return &TreeNode{
		title:    title,
		desc:     desc,
		parent:   parent,
		children: []*TreeNode{},
	}
}
