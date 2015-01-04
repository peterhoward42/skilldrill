package model

// The skillsSerializer type is capable of serialising a list of all the skills
// present in the Api model.
type skillsSerializer struct {
    outBuf *[]byte
    api *Api
}

// Compulsory constructor.
func newSkillsSerializer(outBuf *[]byte, api *Api) *skillsSerializer {
	return &skillsSerializer{outBuf, api}
}
