package model

type skillHoldings struct {
	skillsOfPerson  map[*person][]*treeNode
	peopleWithSkill map[*treeNode][]*person
}

func newSkillHoldings() *skillHoldings {
	return &skillHoldings{
		skillsOfPerson:  map[*person][]*treeNode{},
		peopleWithSkill: map[*treeNode][]*person{}}
}
