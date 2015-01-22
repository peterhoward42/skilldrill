package sets

// The SetOfString type provides the conventional SET model for strings.
type SetOfString struct {
	data map[string]bool
}

// Compulsory constructor.
func NewSetOfString() *SetOfString {
	return &SetOfString{data: map[string]bool{}}
}

// The function Overwrite() replaces the set's content with that of the
// given slice.
func (s *SetOfString) Overwrite(newContent []string) {
	s.data = map[string]bool{}
	for _, value := range newContent {
		s.Add(value)
	}
}

// The function Add(), adds the given value into the set.
func (set *SetOfString) Add(str string) {
	set.data[str] = true
}

// The function Remove(), removes the given value from the set.
func (set *SetOfString) Remove(str string) {
	delete(set.data, str)
}

// The function RemoveIfPresent(), removes the given value from the set if the
// value is present, but does not object when it is not.
func (set *SetOfString) RemoveIfPresent(val string) {
	if set.Contains(val) == false {
		return
	}
	set.Remove(val)
}

// The function Contains() tests for the presence of the given value in
// the set.
func (set *SetOfString) Contains(str string) bool {
	_, ok := set.data[str]
	return ok
}

func (s *SetOfString) MarshalYAML() (interface{}, error) {
	return s.AsSlice(), nil
}

func (s *SetOfString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	tmpSlice := make([]string, 0)
	err := unmarshal(&tmpSlice)
	s.Overwrite(tmpSlice)
	return err
}

func (set *SetOfString) AsSlice() (slice []string) {
	for k, _ := range set.data {
		slice = append(slice, k)
	}
	return slice
}
