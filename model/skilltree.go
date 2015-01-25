package model

type skillNode struct {
    title   string
    desc string
    uid   int
    parent *skillNode
    children []*skillNode
}

type skillTree struct {
    head *skillNode
    nodeFromUid map[int]*skillNode
}

func newSkillTree() (*skillTree) {
    sktr := &skillTree{
        head: nil,
        nodeFromUid: map[int]*skillNode{},
        }
    return sktr
}

func (tree *skillTree) addSkillNode(title string, desc string, 
    parent int) (skillId int, skillNode *skillNode) {
    skillId = tree.nextUid()
    incomer = &skillNode{title, desc, uid, nil, []*skillNode{}}
    tree.nodeFromUid[skillId] = incomer
    if parent == -1 {
        tree.head = incomer
        return
    }
    parentNode := tree.nodesFromUid[parent]
    parentNode.addChild(incomer)
    return skillId, incomer
}
