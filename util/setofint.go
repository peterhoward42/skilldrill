package util

// The SetOfInt type provides the conventional SET model for integers.
type SetOfInt struct {
	data map[int32]bool
}

// Compulsory constructor.
func NewSetOfInt() *SetOfInt {
	return &SetOfInt{data: map[int32]bool{}}
}

func (s *SetOfInt) Contains(val int32) bool {
	_, ok := s.data[val]
	return ok
}

func (s *SetOfInt) Add(val int32) {
	s.data[val] = true
}

func (s *SetOfInt) MarshalYAML() (interface{}, error) {
	return s.AsSlice(), nil
}

func (set *SetOfInt) AsSlice() (slice []int32) {
	for k, _ := range set.data {
		slice = append(slice, k)
	}
	return slice
}
