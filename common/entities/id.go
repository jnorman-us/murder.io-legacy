package entities

type ID struct {
	ID int
}

func (i *ID) GetID() int {
	return i.ID
}

func (i *ID) SetID(id int) {
	i.ID = id
}
