/*
The package sets, provides types to model a set of integers, a set of strings
etc. In addition to the core operations of adding members and testing for the
presence of a given member, the sets implement the yaml package's Marshaler and
UnMarshaller interfaces - so that sets may conveniently be serialized.
*/
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

// The function Remove(), removes the given value from the set.
func (set *SetOfInt) Remove(val int) {
	delete(set.data, val)
}

// The function RemoveIfPresent(), removes the given value from the set if the
// value is present, but does not object when it is not.
func (set *SetOfInt) RemoveIfPresent(val int) {
	if set.Contains(val) == false {
		return
	}
	set.Remove(val)
}

// The function TogglePresenceOf(), removes the given item if it is present,
// or (inversely) adds it, if it is not.
func (set *SetOfInt) TogglePresenceOf(val int) {
	if set.Contains(val) {
		set.Remove(val)
	} else {
		set.Add(val)
	}
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
