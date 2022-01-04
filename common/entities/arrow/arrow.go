package arrow

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

type Arrow struct {
	entities.ID
	collider.Collider

	charge  float64
	shooter *Shooter
}

func NewArrow(s *Shooter, charge float64) *Arrow {
	var shooter = *s
	var arrow = &Arrow{
		shooter: s,
		charge:  charge,
	}
	arrow.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(0, 0), 0, 20, 4),
		},
		[]collider.Circle{},
		1,
	)
	arrow.SetAngle(shooter.GetAngle())
	arrow.SetPosition(shooter.GetPosition())
	//arrow.SetVelocity(shooter.GetVelocity())

	var force = types.NewVector(20, 20)
	force.Rotate(shooter.GetAngle())
	force.Scale(charge)
	arrow.ApplyForce(force)

	return arrow
}

func (a *Arrow) Tick() {

}
