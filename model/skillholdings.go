package model

type skillHoldings struct {
	skillsOfPerson  map[*person][]*skillNode
	peopleWithSkill map[*skillNode][]*person
}

func newSkillHoldings() *skillHoldings {
	return &skillHoldings{
		skillsOfPerson:  map[*person][]*skillNode{},
		peopleWithSkill: map[*skillNode][]*person{}}
}
