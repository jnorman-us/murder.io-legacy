package entities

type ID struct {
	id int
}

func (i *ID) GetID() int {
	return i.id
}

func (i *ID) SetID(id int) {
	i.id = id
}
