package model

// The setOfSkills type provides the contentional SET model for the type
// skillNode.
type setOfSkills struct {
	data map[*skillNode]bool
}

// Compulsory constructor.
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
