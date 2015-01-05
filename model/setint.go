package model

// The setOfInt type provides the contentional SET model for integers.
type setOfInt struct {
	data map[int32]bool
}

// Compulsory constructor.
func newSetOfInt() *setOfInt {
	return &setOfInt{data: map[int32]bool{}}
}

func (s *setOfInt) contains(val int32) bool {
	_, ok := s.data[val]
	return ok
}

func (s *setOfInt) add(val int32) {
	s.data[val] = true
}
