package model

type SkillHoldings struct {
	skillsOfPerson  map[*Person][]*TreeNode
	peopleWithSkill map[*TreeNode][]*Person
}

func NewSkillHoldings() *SkillHoldings {
	return &SkillHoldings{
		skillsOfPerson:  map[*Person][]*TreeNode{},
		peopleWithSkill: map[*TreeNode][]*Person{}}
}
