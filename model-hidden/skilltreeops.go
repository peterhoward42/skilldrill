package model

import (
	"github.com/peterhoward42/skilldrill/util/sets"
	"strings"
)

/*
The skillTreeOps type is a place for algorithmic functions to live that depend
on traversing parent child relationships in the skills taxonomy tree. The aim
is to prevent any other parts of the model software from having to engage with
this topic.
*/
type skillTreeOps struct {
	api *Api
}

/*
The skillWording() method is capable of assembling a description for a
skillNode that is based on its child-parent ancestry. In other words it can
synthesise a description that describes the skill completely and autonomuously,
by contatenating the skill node descriptions, working up the tree. It provides
for convenience also the skill Node title and the aggregate description broken
into pieces. The <desc> return value is from the leaf node alone. The
<descInContext> return value is the aggregated description. While the
<contextAlone> is the description drawn from the skill node's parent
(recursively).
*/
func (treeOps *skillTreeOps) skillWording(skill *skillNode) (title string,
	desc string, descInContext string, contextAlone string) {
	nodes := []*skillNode{}
	treeOps.lineageOf(skill, &nodes)
	descriptions := []string{}
	for _, node := range nodes {
		descriptions = append(descriptions, node.Desc)
	}
	title = skill.Title
	desc = skill.Desc
	descInContext = strings.Join(descriptions, ">>>")
	contextAlone = strings.Join(descriptions[:len(descriptions)-1], ">>>")
	return
}

/*
The lineageOf() method provides the list of skillNodes that comprise the
parent chain from the given skill up to the root of the tree. Root first.
*/
func (treeOps *skillTreeOps) lineageOf(skill *skillNode,
	lineage *[]*skillNode) {
	// recurse to add parent lineage first
	if skill.Parent != -1 {
		treeOps.lineageOf(treeOps.api.skillFromId[skill.Parent], lineage)
	}
	// now add me
	*lineage = append(*lineage, skill)
}

/*
The method enumerateTree() provides a list of skill Uids in the order they
should appear when displaying the tree. It is person-specific, and omits the
nodes that have been collapsed (using CollapseSkill()) - including their
children.
*/
func (treeOps *skillTreeOps) enumerateTree(collapsedNodes *sets.SetOfInt) (
	skills []int, depths []int) {
	curNode := treeOps.api.skillFromId[treeOps.api.SkillRoot]
	skills = []int{}
	depths = []int{}
	curDepth := 0
	// Recursive
	treeOps.enumerateNode(curNode, collapsedNodes, curDepth, &skills, &depths)
	return
}

// Recursive helper for EnumerateTree() method.
func (treeOps *skillTreeOps) enumerateNode(curNode *skillNode,
	collapsedNodes *sets.SetOfInt, curDepth int, skills *[]int, depths *[]int) {
	// Me first
	*skills = append(*skills, curNode.Uid)
	*depths = append(*depths, curDepth)

	// If I am collapsed, do not continue to recurse into my children
	if collapsedNodes.Contains(curNode.Uid) {
		return
	}
	// Otherwise, do
	childDepth := curDepth + 1
	for _, child := range curNode.Children {
		treeOps.enumerateNode(treeOps.api.skillFromId[child], collapsedNodes,
			childDepth, skills, depths)
	}
	return
}
