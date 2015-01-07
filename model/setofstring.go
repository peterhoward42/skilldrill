package model

// The setOfString type provides the conventional SET model for strings.
type setOfString struct {
	data map[string]bool
}

// Compulsory constructor.
func newSetOfString() *setOfString {
	return &setOfString{data: map[string]bool{}}
}

func (set *setOfString) contains(str string) bool {
	_, ok := set.data[str]
	return ok
}

func (set *setOfString) add(str string) {
	set.data[str] = true
}

func (s *setOfString) MarshalYAML() (interface{}, error) {
	return s.AsSlice(), nil
}

func (set *setOfString) AsSlice() (slice []string) {
	for k, _ := range set.data {
		slice = append(slice, k)
	}
	return slice
}
