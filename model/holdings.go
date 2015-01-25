package model

type holdings struct {
    skillsOfPeople map[string][]*skillNode
    peopleWithSkill map[*skillNode][]*string
}

func newHoldings() *holdings {
    return &holdings{
        skillsOfPeople: make(map[string][]*skillNode),
        peopleWithSkill: make(map[*skillNode][]*string),
    }
}
