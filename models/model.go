package models

type Model inteface {
    GetId() string
    SetId(id string)
}

type modelImpl struct {
    id string
}

func (m *ModelImpl) SetId(id string) {
    m.id = id
}
