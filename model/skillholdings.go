package model

import (
	"github.com/peterhoward42/skilldrill/util/sets"
)

/*
The skillHoldings type contains bindings between people and the set of skills
they hold.  The design intent is that none of fields are exported, but the
reason that some are, is solely to facilitate automated serialization by
yaml.Marshal().
*/
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

/*
The registerSkill() method makes the given skill uid known to the object. It is
harmless to call it when the skill has already been registered.
*/
func (sh *skillHoldings) registerSkill(skillId int) {
	if _, ok := sh.PeopleWithSkill[skillId]; ok {
		return
	}
	sh.PeopleWithSkill[skillId] = sets.NewSetOfString()
}

// The registerPerson() method makes the given person email known to the object.
// It is harmless to call it when the email has already been registered.
func (sh *skillHoldings) registerPerson(email string) {
	if _, ok := sh.SkillsOfPerson[email]; ok {
		return
	}
	sh.SkillsOfPerson[email] = sets.NewSetOfInt()
}

/*
The UnRegisterPerson method, removes all traces of the given person from the
data that this class holds.
*/
func (sh *skillHoldings) UnRegisterPerson(toGo person) {
	setOfSkills := sh.SkillsOfPerson[toGo.Email]
	for _, skillId := range setOfSkills.AsSlice() {
		sh.PeopleWithSkill[skillId].Remove(toGo.Email)
	}
	delete(sh.SkillsOfPerson, toGo.Email)
}

/*
The method bind() adds the given skill to the set of skills held for the given
person. The skill and the person are automatically registered if they have not
been previously.  An error is generated if the skill is a CATEGORY.
*/
func (sh *skillHoldings) bind(skill int, person string) {
	sh.registerSkill(skill)
	sh.registerPerson(person)

	sh.SkillsOfPerson[person].Add(skill)
	sh.PeopleWithSkill[skill].Add(person)
}
