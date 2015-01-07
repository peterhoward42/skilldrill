package model

import (
)

// The skillHoldings type contains bindings between people and the set of skills
// they hold.
// The design intent is that none of fields are exported, but the reason
// that some are, is solely to facilitate automated serialization by
// yaml.Marshal().
type skillHoldings struct {
	SkillsOfPerson  map[string]*setOfInt   // email -> skill.Uid
	PeopleWithSkill map[int32]*setOfString // skill.Uid -> email
}

// Compulsory constructor.
func newSkillHoldings() *skillHoldings {
	return &skillHoldings{
		SkillsOfPerson:  map[string]*setOfInt{},
		PeopleWithSkill: map[int32]*setOfString{},
	}
}

// The method bind() adds the given skill to the set of skills held for the given
// person. An error is generated if the skill is a CATEGORY.
func (sh *skillHoldings) bind(skill int32, email string) {
	skills, ok := sh.SkillsOfPerson[email]
	if !ok {
		skills = newSetOfInt()
		sh.SkillsOfPerson[email] = skills
	} else {
	}

	skills.add(skill)

	people, ok := sh.PeopleWithSkill[skill]
	if !ok {
		people = newSetOfString()
		sh.PeopleWithSkill[skill] = people
	}
	people.add(email)
}

func (holdings *skillHoldings) MarshalYAML() (interface{}, error) {
	d := make(map[string]interface{})
	d["skillsOfPerson"] = holdings.yamlSkillsOfPerson()
	d["peopleWithSkill"] = holdings.yamlPeopleWithSkill()
	return d, nil
}

func (holdings *skillHoldings) yamlSkillsOfPerson() interface{} {
	d := make(map[string][]int32)
	for email, skills := range holdings.SkillsOfPerson {
		sas := skills.asSlice()
		d[email] = sas
	}
	return d
}

func (holdings *skillHoldings) yamlPeopleWithSkill() interface{} {
	d := make(map[int32][]string)
	for skillUid, people := range holdings.PeopleWithSkill {
		d[skillUid] = people.asSlice()
	}
	return d
}
