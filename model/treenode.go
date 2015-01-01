package model

type TreeNode struct {
	title    string
	desc     string
	parent   *TreeNode
	children []*TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{children: []*TreeNode{}}
}
