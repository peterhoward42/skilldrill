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


func (holdings *holdings) personAdded(emailName string) {
    holdings.skillsOfPeople[emailName] = []*skillNode{}
}
