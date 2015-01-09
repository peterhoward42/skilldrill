package sets

// The SetOfInt type provides the conventional SET model for integers.
type SetOfInt struct {
	data map[int]bool
}

// Compulsory constructor.
func NewSetOfInt() *SetOfInt {
	return &SetOfInt{data: map[int]bool{}}
}

// The function Overwrite() replaces the set's content with that of the
// given slice.
func (s *SetOfInt) Overwrite(newContent []int) {
	s.data = map[int]bool{}
	for _, value := range newContent {
		s.Add(value)
	}
}

// The function Add(), adds the given value into the set.
func (s *SetOfInt) Add(val int) {
	s.data[val] = true
}

// The function Contains() tests for the presence of the given value in
// the set.
func (s *SetOfInt) Contains(val int) bool {
	_, ok := s.data[val]
	return ok
}

func (s *SetOfInt) MarshalYAML() (interface{}, error) {
	return s.AsSlice(), nil
}

func (s *SetOfInt) UnmarshalYAML(unmarshal func(interface{}) error) error {
	tmpSlice := make([]int, 0)
	err := unmarshal(&tmpSlice)
	s.Overwrite(tmpSlice)
	return err
}

// The function AsSlice() provides a slice of integers comprising the
// members of the set.
func (set *SetOfInt) AsSlice() (slice []int) {
	for k, _ := range set.data {
		slice = append(slice, k)
	}
	return slice
}
