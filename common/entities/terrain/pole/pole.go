package pole

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Pole struct {
	types.ID
	collider.Collider
}

func NewPole() *Pole {
	var pole = &Pole{}
	pole.Setup()
	return pole
}

func (p *Pole) Setup() {
	p.Collider.SetupCollider(
		[]collider.Rectangle{},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, 0), 10),
		},
	)
}
