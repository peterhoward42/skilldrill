package model

type skillHoldings struct {
	skillsOfPerson  map[*person]*setOfSkills
	peopleWithSkill map[*skillNode]*setOfPeople
}

func newSkillHoldings() *skillHoldings {
	return &skillHoldings{
		skillsOfPerson:  map[*person]*setOfSkills{},
		peopleWithSkill: map[*skillNode]*setOfPeople{},
	}
}

func (skh *skillHoldings) bind(skill *skillNode, personWithSkill *person) {
	skills, ok := skh.skillsOfPerson[personWithSkill]
	if !ok {
		skills = newSetOfSkills()
		skh.skillsOfPerson[personWithSkill] = skills
	}
	skills.add(skill)

	people, ok := skh.peopleWithSkill[skill]
	if !ok {
		people = newSetOfPeople()
		skh.peopleWithSkill[skill] = people
	}
	people.add(personWithSkill)
}
