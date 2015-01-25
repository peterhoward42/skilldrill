package model

type skillNode struct {
    title   string
    desc string
    uid   int
}

type skillTree struct {
    head *skillNode
    nodeFromUid map[int]*skillNode
    parent map[*skillNode]*skillNode
    children map[*skillNode][]*skillNode // slices are kept sorted by title
}

func newSkillTree() (*skillTree) {
    sktr := &skillTree{
        head: nil,
        nodeFromUid: make(map[int]*skillNode),
        parent : make(map[*skillNode]*skillNode),
        children: make(map[*skillNode][]*skillNode),
        }
    return sktr
}
