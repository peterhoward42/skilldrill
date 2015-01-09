package model

import (
	"github.com/peterhoward42/skilldrill/util/sets"
)

// The skillHoldings type contains bindings between people and the set of skills
// they hold.
// The design intent is that none of fields are exported, but the reason
// that some are, is solely to facilitate automated serialization by
// yaml.Marshal().
type skillHoldings struct {
	SkillsOfPerson  map[string]*sets.SetOfInt // email -> skill.Uid
	PeopleWithSkill map[int]*sets.SetOfString // skill.Uid -> email
}

// Compulsory constructor.
func newSkillHoldings() *skillHoldings {
	return &skillHoldings{
		SkillsOfPerson:  map[string]*sets.SetOfInt{},
		PeopleWithSkill: map[int]*sets.SetOfString{},
	}
}

// The method bind() adds the given skill to the set of skills held for the given
// person. An error is generated if the skill is a CATEGORY.
func (sh *skillHoldings) bind(skill int, email string) {
	skills, ok := sh.SkillsOfPerson[email]
	if !ok {
		skills = sets.NewSetOfInt()
		sh.SkillsOfPerson[email] = skills
	} else {
	}

	skills.Add(skill)

	people, ok := sh.PeopleWithSkill[skill]
	if !ok {
		people = sets.NewSetOfString()
		sh.PeopleWithSkill[skill] = people
	}
	people.Add(email)
}
