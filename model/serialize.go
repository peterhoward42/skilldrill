package model

// The yamlModel type is an in-memory representation of the skilldrill model
// that is amenable to automated serialization into yaml. It represents relations
// using only UIDs (not pointers), and exposes publicly the fields it wants
// yaml.Marshal() to discover.
type yamlModel struct {
	api        *Api
	SkillsList []*yamlSkill
}

// Compulsory constructor.
func newYamlModel(api *Api) (m *yamlModel) {
	m = &yamlModel{api: api}
	m.populateSkills()
	return
}

type yamlSkill struct {
	Uid      int64
	Role     string
	Title    string
	Desc     string
	Parent   int64
	Children []int64
}

func (m *yamlModel) populateSkills() {
	for _, skill := range m.api.skillFromId {
		m.SkillsList = append(m.SkillsList, &yamlSkill{
			skill.uid,
			skill.role,
			skill.title,
			skill.desc,
			parentUid(skill),
			skill.childUids(),
		})
	}
	return
}

func parentUid(skill *skillNode) (uid int64) {
	if skill.parent == nil {
		uid = -1
        return
	}
	uid = skill.parent.uid
    return
}
