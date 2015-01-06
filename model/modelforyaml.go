package model

/*
The modelForYaml type is a data structure derived from model.Api that contains
only those parts that must be serialized, and which exposes publically some
fields for yaml.Marshal() to discover. It contains pointers into the Api model
and hence, is not independent from it from a concurrency point of view.
*/

type modelForYaml struct {
}

// Compulsory constructor.
func NewModelForYaml(api *Api) (m *modelForYaml) {
    m = &modelForYaml{}
    return
}
