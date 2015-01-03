package model

// The setOfPeople type provides the contentional SET model for the type person.
type setOfPeople struct {
	data map[*person]bool
}

// Compulsory constructor.
func newSetOfPeople() *setOfPeople {
	return &setOfPeople{data: map[*person]bool{}}
}

func (s *setOfPeople) contains(thisPerson *person) bool {
	_, ok := s.data[thisPerson]
	return ok
}

func (s *setOfPeople) add(thisPerson *person) {
	s.data[thisPerson] = true
}
