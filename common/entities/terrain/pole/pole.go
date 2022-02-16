package pole

import (
	collider2 "github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Pole struct {
	types.ID
	collider2.Collider
}

func NewPole() *Pole {
	var pole = &Pole{}
	pole.Setup()
	return pole
}

func (p *Pole) Setup() {
	p.Collider.SetupCollider(
		[]collider2.Rectangle{},
		[]collider2.Circle{
			collider2.NewCircle(types.NewVector(0, 0), 10),
		},
	)
}
