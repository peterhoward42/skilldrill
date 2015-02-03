package model

import (
	"github.com/peterhoward42/skilldrill/util/sets"
)

/*
The holdings type holds which people have which skill and offers CRUD
operations on this data.
*/
type holdings struct {
	skillsOfPeople  map[string]*sets.SetOfInt
	peopleWithSkill map[*skillNode]*sets.SetOfString
}

func newHoldings() *holdings {
	return &holdings{
		skillsOfPeople:  map[string]*sets.SetOfInt{},
		peopleWithSkill: map[*skillNode]*sets.SetOfString{},
	}
}

func (holdings *holdings) notifyPersonAdded(emailName string) {
	holdings.skillsOfPeople[emailName] = sets.NewSetOfInt()
}

func (holdings *holdings) notifySkillAdded(incomer *skillNode) {
	holdings.peopleWithSkill[incomer] = sets.NewSetOfString()
}

func (holdings *holdings) personExists(emailName string) bool {
	_, exists := holdings.skillsOfPeople[emailName]
	return exists
}

func (holdings *holdings) skillExists(skillNode *skillNode) bool {
	_, exists := holdings.peopleWithSkill[skillNode]
	return exists
}

func (holdings *holdings) personHasSkill(skillNode *skillNode, email string) (
    hasSkill bool) {
    return holdings.skillsOfPeople[email].Contains(skillNode.uid)
}

func (holdings *holdings) givePersonSkill(skill *skillNode, emailName string) {
	holdings.skillsOfPeople[emailName].Add(skill.uid)
	holdings.peopleWithSkill[skill].Add(emailName)
}

func (holdings *holdings) someoneHasThisSkill(skill *skillNode) bool {
	return len(holdings.peopleWithSkill[skill].AsSlice()) != 0
}
