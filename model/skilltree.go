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
        title: title,
        desc: desc,
        uid: uid,
        parent: parent,
    }
}

type skillTree struct {
	root        *skillNode
	nodeFromUid map[int]*skillNode
    nextId int
}

func newSkillTree() *skillTree {
	return  &skillTree{
		root:        nil,
		nodeFromUid: map[int]*skillNode{},
        nextId : 0,
	}
}

// To create a root node, call this with parent set to -1
func (tree *skillTree) addSkillNode(title string, desc string,
	parent int) (skillId int, newNode *skillNode) {

    skillId = tree.nextUid()

    // code block partly duplicated for clarity
    if parent == -1 {
        newNode = newSkillNode(title, desc, skillId, nil)
        tree.nodeFromUid[skillId] = newNode
        tree.root = newNode
    } else {
        parentSkill := tree.nodeFromUid[parent]
        newNode = newSkillNode(title, desc, skillId, parentSkill)
        tree.nodeFromUid[skillId] = newNode
        parentSkill.addChild(newNode)
    }
    return
}

func (tree *skillTree) nextUid() int {
    tree.nextId += 1
    return tree.nextId
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

