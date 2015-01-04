package model

// The yamlModel type is an in-memory representation of the skilldrill model 
// that is amenable to automated serialization into yaml. It represents relations
// using only UIDs (not pointers), and exposes the fields it wants yaml.Marshal()
// to discover.
type yamlModel struct {
    api *Api
    SkillsList *[]yamlSkill
}

// Compulsory constructor.
func newYamlModel(api *Api) (m *yamlModel) {
    m = &yamlModel{
        api: api,
        SkillsList: &[]yamlSkill{},
    }
    m.populateSkills()
    return
}

type yamlSkill struct {
    Uid int64
    Role string
    Title string
    Desc string
    Parent int64
    Children []int64
}

func (m *yamlModel) populateSkills() {
    for _, skill := range(m.api.skillFromId) {
        ySkill := &yamlSkill{
            skill.uid,
            skill.role,
            skill.title,
            skill.desc,
            skill.parent.uid,
            skill.childUids(),
        }
        append(m.SkillsList, ySkill)
    }
    return
}
