package model

// The skillHoldings type contains bindings between people and the set of skills
// they hold.
type skillHoldings struct {
	skillsOfPerson  map[*person]*setOfSkills
	peopleWithSkill map[*skillNode]*setOfPeople
}

// Compulsory constructor.
func newSkillHoldings() *skillHoldings {
	return &skillHoldings{
		skillsOfPerson:  map[*person]*setOfSkills{},
		peopleWithSkill: map[*skillNode]*setOfPeople{},
	}
}

// The method bind() adds the given skill to the set of skills held for the given
// person. An error is generated if the skill is a CATEGORY.
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
