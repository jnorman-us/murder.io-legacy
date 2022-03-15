package types

type ID int32

func (i ID) GetID() ID {
	return i
}
