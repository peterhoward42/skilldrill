package util

// The SetOfString type provides the conventional SET model for strings.
type SetOfString struct {
	data map[string]bool
}

// Compulsory constructor.
func NewSetOfString() *SetOfString {
	return &SetOfString{data: map[string]bool{}}
}

func (set *SetOfString) Contains(str string) bool {
	_, ok := set.data[str]
	return ok
}

func (set *SetOfString) Add(str string) {
	set.data[str] = true
}

func (s *SetOfString) MarshalYAML() (interface{}, error) {
	return s.AsSlice(), nil
}

func (set *SetOfString) AsSlice() (slice []string) {
	for k, _ := range set.data {
		slice = append(slice, k)
	}
	return slice
}
