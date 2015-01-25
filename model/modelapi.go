package model

type ModelApi struct {
    impl *modelImpl
}

func NewModelApi() *ModelApi {
    return &ModelApi{
        impl: newModelImpl(),
    }
}

func (api *ModelApi) AddPerson(emailName string) {
    api.impl.addPerson(emailName)
}

func (api *ModelApi) AddSkillNode(title string, description string, 
    parent int) (uid int) {
    return api.impl.addSkillNode(title, description, parent)
}
