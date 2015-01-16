package model

import (
	"sort"
)

/*
The skillNode type models a node in the skill hierachy in terms of its title,
description and parent child relations. Only the leaf-level nodes are permitted
to be SKILLS. The others must have the role CATEGORY. The descriptive strings
should provide only a qualification for their specialism with respect to their
parent category, and should not duplicate this information.  The design intent
is that none of Api fields are exported, but the reason that some are, is
solely to facilitate automated serialization by yaml.Marshal(). The node's
children are maintained in alphabetical order by title, and to support this
behaviour, the caller must dependency-inject to the constructor, a mapper of
skillId to skill title. This avoids having to duplicate in the node information
about the world outside of itself.
*/
type skillNode struct {
	Uid      int
	Role     string // SKILL | CATEGORY
	Title    string
	Desc     string
	Parent   int
	Children []int // Do not alter this directly, use addChild()
	mapper   titleMapper
}

// Compulsory constructor.
func newSkillNode(uid int, role string, title string, desc string,
	parent int, mapper titleMapper) *skillNode {
	return &skillNode{
		Uid:      uid,
		Role:     role,
		Title:    title,
		Desc:     desc,
		Parent:   parent,
		Children: []int{},
		mapper:   mapper,
	}
}

/*
The method addChild() adds the given skill uid into the list held of this
node's children - whilst maintaining their alphabetical order. Not completely
straightforward, because the skill node knows only about the Uid's of the other
children, and not (in of itself) their titles.
*/
func (skill *skillNode) addChild(newChild int) {
	skill.Children = append(skill.Children, newChild)

	// Reestablish the order by obtaining the titles of the peer children and
	// noting the consequences of sorting the list of titles.

	uids := map[string]int{} // uids keyed on titles
	titles := []string{}
	for _, uid := range skill.Children {
		title := skill.mapper.titleFromId(uid)
		uids[title] = uid
		titles = append(titles, title)
	}
	sort.Strings(titles)
	for idx, title := range titles {
		skill.Children[idx] = uids[title]
	}
}

/*
The method removeChild() removes the given skill uid from the list held of this
node's children - whilst maintaining their alphabetical order.
*/
func (skill *skillNode) removeChild(toRemove int) {
	replacement := []int{}
	for _, uid := range skill.Children {
		if uid != toRemove {
			replacement = append(replacement, uid)
		}
	}
	skill.Children = replacement
}

type titleMapper interface {
	titleFromId(skillUid int) (title string)
}
