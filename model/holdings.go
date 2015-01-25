package model

import (
	"github.com/peterhoward42/skilldrill/util/sets"
)

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
