package model

// The setOfInt type provides the conventional SET model for integers.
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

func (set *setOfInt) asSlice() (slice []int32) {
	for k, _ := range set.data {
		slice = append(slice, k)
	}
	return slice
}
