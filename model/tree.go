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

func (tree *skillTree) addRootSkillNode(title string,
	desc string) (skillId int, newNode *skillNode) {
	skillId = tree.nextUid()
	newNode = newSkillNode(title, desc, skillId, nil)
	tree.nodeFromUid[skillId] = newNode
	tree.root = newNode
}

func (tree *skillTree) addChildSkillNode(title string, desc string,
	parent int) (skillId int, newNode *skillNode) {
	skillId = tree.nextUid()
	parentSkill := tree.nodeFromUid[parent]
	newNode = newSkillNode(title, desc, skillId, parentSkill)
	tree.nodeFromUid[skillId] = newNode
	parentSkill.addChild(newNode)
}

//----------------------------------------------------------------------------
// Query operations
//----------------------------------------------------------------------------

func (tree *skillTree) treeIsEmpty() bool {
    return tree.root == nil
}

//----------------------------------------------------------------------------
// Internal operations
//----------------------------------------------------------------------------

func (tree *skillTree) nextUid() int {
	tree.nextId += 1
	return tree.nextId
}
