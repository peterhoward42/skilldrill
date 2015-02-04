package model

/*
The tree type owns the storage of the skill nodes and their tree-like topology.
The skills themselves are modelled by the closely-couple skillNode type, and
each skill node individually contains references to its parent and children.
However the only object that knows which is the skill at the root of the tree
is this one, and all the logic that deals with the tree-like topology belongs
in this tree type.
*/
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

func (tree *tree) titleOfSkill(skillId int) (title string) {
	return tree.nodeFromUid[skillId].title
}

/*
The EnumerateTree method provides a linear sequence of the skill Uids which
can be used essentiall as an iteratorto used to render the skill tree. Separate
query methods are available to get the extra data that might be needed for
each row - like for example its depth in the tree. It is personalised to a
given emailName, and will have omitted the children of any skill nodes the
person has collapsed.  Errors: UnknownPerson
*/
func (tree *treel) enumerateTree(emailName string) (treeRows []int) {
	return tree.addRowsRecursively(&treeRows, tree.root)
}

//----------------------------------------------------------------------------
// Internal operations
//----------------------------------------------------------------------------

func (tree *tree) nextUid() int {
	tree.nextId += 1
	return tree.nextId
}


func (tree *tree) addRowsRecursively(treeRows *[]int, startNode *skillNode) {
    *treeRows = append(*treeRows, startNode.skillUid)
    for _, child := range startNode.children {
        tree.addRowsRecursively(treeRows, child)
    }
}
