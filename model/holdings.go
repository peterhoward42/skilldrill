package model

type holdings struct {
    skillsOfPeople map[string][]*skillNode
    peopleWithSkill map[*skillNode][]*string
}

func newHoldings() *holdings {
    return &holdings{
        skillsOfPeople: map[string][]*skillNode{},
        peopleWithSkill: map[*skillNode][]*string{},
    }
}


func (holdings *holdings) addPerson(emailName string) {
    holdings.skillsOfPeople[emailName] = []*skillNode{}
}

func (holdings *holdings) notifySkillAdded(incomer *skillNode) {
    holdings.peopleWithSkill[incomer] = []*string{}
}

