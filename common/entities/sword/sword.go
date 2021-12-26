package sword

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

type Sword struct {
	entities.ID
	collider.Collider
}

func NewSword() *Sword {
	var sword = &Sword{}
	sword.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(21, 0), 0, 18, 2),
		},
		[]collider.Circle{},
		1,
	)
	return sword
}

func (s *Sword) Tick() {

}
