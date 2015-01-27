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

/*
The EnumerateTree method provides a linear sequence of TreeDisplayItem which
can be used to used to render the skill tree. It is personalised to a given
emailName, and will have omitted the children of any skill nodes the person has
collapsed.
*/
func (tree *tree) EnumerateTree(collapsedNodes []*skillNode) (
    displayRows []TreeDisplayItem) {
    displayRows = []TreeDisplayItem{}
    // Use recursive walker
    depth := 0
    addRowsRecursively(tree.root, depth, &displayRows)
    return
}
`

//----------------------------------------------------------------------------
// Internal operations
//----------------------------------------------------------------------------

func (tree *tree) nextUid() int {
	tree.nextId += 1
	return tree.nextId
}

func (tree *tree) addRowsRecursively(startNode *skillNode, 
    depth int, rows *[]TreeDisplayItem) {
    *rows = append(*rows, tree.makeOneRow(startNode, depth)
    newDepth := depth + 1
    for _, child := range startNode.children {
        tree.addRowsRecursively(child, newDepth, rows)
    }
}

func (tree *tree) makeOneRow(skill *skillNode, depth int) (
    row *TreeDisplayItem) {
    fart got to here
}
