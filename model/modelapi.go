package model

type ModelApi struct {
    impl *modelImpl
}

func NewModelApi() *ModelApi {
    return &ModelApi{
        impl: newModelImpl(),
    }
}
