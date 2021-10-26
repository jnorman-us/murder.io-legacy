package classes

type Identifiable interface {
	SetID(id int32)
	GetID() int32
	Tick()
}
