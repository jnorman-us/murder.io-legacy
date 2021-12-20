package innocent

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Innocent struct {
	id      int
	spawner Spawner
	collider.Collider
}

func NewInnocent() *Innocent {
	var innocent = &Innocent{}
	innocent.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(0, -20), .3, 20, 20),
		},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, 20), 10),
		},
	)
	return innocent
}

func (i *Innocent) GetID() int {
	return i.id
}

func (i *Innocent) SetID(id int) {
	i.id = id
}
