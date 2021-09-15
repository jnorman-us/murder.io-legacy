package entities

import (
	"github.com/josephnormandev/murder/collider"
	world2 "github.com/josephnormandev/murder/world"
)

type Entity struct {
	collider.Collider
	id    int32
	world *world2.World
}

func (e *Entity) GetID() int32 {
	return e.id
}

func (e *Entity) SetID(id int32) {
	e.id = id
}
