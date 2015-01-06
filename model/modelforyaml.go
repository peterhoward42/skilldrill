package model

/*
The modelForYaml type is a data structure derived from model.Api that contains
only those parts that must be serialized, and which exposes publicly some
fields for yaml.Marshal() to discover. It contains pointers into the Api model
and hence, is not independent from it from a concurrency point of view.
*/

type modelForYaml struct {
	api           *Api           // not serialized
	Skills        []*skillNode   // flat list of skills (serialized)
	People        []*person      // flast of people (serialized)
	SkillHoldings *skillHoldings // delegate to the skillHoldings type
}

// Compulsory constructor.
func NewModelForYaml(api *Api) (m *modelForYaml) {
	m = &modelForYaml{
		api:           api,
		Skills:        listOfSkills(api),
		People:        listOfPeople(api),
		SkillHoldings: api.skillHoldings,
	}
	return
}

func listOfSkills(api *Api) (theList []*skillNode) {
	for _, skill := range api.skills {
		theList = append(theList, skill)
	}
	return theList
}

func listOfPeople(api *Api) (theList []*person) {
	for _, person := range api.people {
		theList = append(theList, person)
	}
	return theList
}
