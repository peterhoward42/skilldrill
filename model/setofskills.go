package model

type setOfSkills struct {
	data map[*skillNode]bool
}

func newSetOfSkills() *setOfSkills {
	return &setOfSkills{data: map[*skillNode]bool{}}
}

func (s *setOfSkills) contains(skill *skillNode) bool {
	_, ok := s.data[skill]
	return ok
}

func (s *setOfSkills) add(skill *skillNode) {
	s.data[skill] = true
}
