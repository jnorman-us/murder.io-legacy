package types

type ID int

func (i ID) GetID() ID {
	return i
}
