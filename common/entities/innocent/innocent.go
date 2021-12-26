package innocent

import (
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
	"math"
)

type Innocent struct {
	entities.ID
	spawner *Spawner
	collider.Collider

	input input.Input
}

func NewInnocent() *Innocent {
	var innocent = &Innocent{}
	innocent.SetupCollider(
		[]collider.Rectangle{},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, 0), 10),
		},
		10,
	)
	return innocent
}

func (i *Innocent) Tick() {
	var in = i.input
	var angle = 0.0
	var movementForce = types.NewVector(20, 0)
	if in.Left && in.Forward {
		angle = math.Pi / 4 * 5
	} else if in.Left && in.Backward {
		angle = math.Pi / 4 * 3
	} else if in.Right && in.Forward {
		angle = math.Pi / 4 * 7
	} else if in.Right && in.Backward {
		angle = math.Pi / 4 * 1
	} else if in.Left {
		angle = math.Pi
	} else if in.Forward {
		angle = math.Pi / 2 * 3
	} else if in.Backward {
		angle = math.Pi / 2
	} else if in.Right {
		angle = 0
	} else {
		movementForce.Scale(0)
	}

	movementForce.RotateAbout(angle, types.NewZeroVector())
	i.Collider.ApplyForce(movementForce)
}
