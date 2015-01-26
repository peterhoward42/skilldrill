package model

type tree struct {
	root        *skillNode
	nodeFromUid map[int]*skillNode
	nextId      int
}

func newTree() *tree {
	return &tree{
		root:        nil,
		nodeFromUid: map[int]*skillNode{},
		nextId:      0,
	}
}

//----------------------------------------------------------------------------
// Add operations
//----------------------------------------------------------------------------

func (tree *tree) addRootSkillNode(title string,
	desc string) (newNode *skillNode, skillId int) {
	skillId = tree.nextUid()
	newNode = newSkillNode(title, desc, skillId, nil)
	tree.nodeFromUid[skillId] = newNode
	tree.root = newNode
	return
}

func (tree *tree) addChildSkillNode(title string, desc string,
	parent int) (newNode *skillNode, skillId int) {
	skillId = tree.nextUid()
	parentSkill := tree.nodeFromUid[parent]
	newNode = newSkillNode(title, desc, skillId, parentSkill)
	tree.nodeFromUid[skillId] = newNode
	parentSkill.addChild(newNode)
	return
}

//----------------------------------------------------------------------------
// Query operations
//----------------------------------------------------------------------------

func (tree *tree) treeIsEmpty() bool {
	return tree.root == nil
}

func (tree *tree) skillExists(skillId int) bool {
	_, exists := tree.nodeFromUid[skillId]
	return exists
}

//----------------------------------------------------------------------------
// Internal operations
//----------------------------------------------------------------------------

func (tree *tree) nextUid() int {
	tree.nextId += 1
	return tree.nextId
}
