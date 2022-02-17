package pole

import (
	"github.com/josephnormandev/murder/common/collider"
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
		map[string]collider.Rectangle{},
		map[string]collider.Circle{
			"t": collider.NewCircle(types.NewVector(0, 0), 10),
		},
	)
	p.Collider.SetColor(types.Colors.Green)
}
