package model

// The skillHoldings type contains bindings between people and the set of skills
// they hold.
type skillHoldings struct {
	skillsOfPerson  map[string]*setOfInt   // email -> skill.Uid
	peopleWithSkill map[int32]*setOfString // skill.Uid -> email
}

// Compulsory constructor.
func newSkillHoldings() *skillHoldings {
	return &skillHoldings{
		skillsOfPerson:  map[string]*setOfInt{},
		peopleWithSkill: map[int32]*setOfString{},
	}
}

// The method bind() adds the given skill to the set of skills held for the given
// person. An error is generated if the skill is a CATEGORY.
func (sh *skillHoldings) bind(skill int32, email string) {
	skills, ok := sh.skillsOfPerson[email]
	if !ok {
		skills = newSetOfInt()
		sh.skillsOfPerson[email] = skills
	}
	skills.add(skill)

	people, ok := sh.peopleWithSkill[skill]
	if !ok {
		people = newSetOfString()
		sh.peopleWithSkill[skill] = people
	}
	people.add(email)
}
