package bounds

import (
	collider2 "github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Bounds struct {
	types.ID
	collider2.Collider
}

func NewBounds() *Bounds {
	var bounds = &Bounds{}
	bounds.Setup()
	return bounds
}

func (b *Bounds) Setup() {
	b.Collider.SetupCollider(
		[]collider2.Rectangle{
			collider2.NewRectangle(types.NewVector(-200, 0), 0, 20, 1000),
			collider2.NewRectangle(types.NewVector(200, 0), 0, 20, 1000),
			collider2.NewRectangle(types.NewVector(0, 500), 0, 20, 400),
			collider2.NewRectangle(types.NewVector(0, -500), 0, 20, 400),
		},
		[]collider2.Circle{},
	)
}
