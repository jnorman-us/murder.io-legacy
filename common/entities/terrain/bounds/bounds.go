package bounds

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Bounds struct {
	types.ID
	collider.Collider
}

func NewBounds() *Bounds {
	var bounds = &Bounds{}
	bounds.Setup()
	return bounds
}

func (b *Bounds) Setup() {
	b.Collider.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(-200, 0), 0, 20, 1000),
			collider.NewRectangle(types.NewVector(200, 0), 0, 20, 1000),
			collider.NewRectangle(types.NewVector(0, 500), 0, 20, 400),
			collider.NewRectangle(types.NewVector(0, -500), 0, 20, 400),
		},
		[]collider.Circle{},
	)
}
