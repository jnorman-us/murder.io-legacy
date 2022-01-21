package arrow

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

type Arrow struct {
	entities.ID
	collider.Collider

	spawner *Spawner
	Charge  float64
	shooter *Shooter
}

func NewArrow(s *Shooter, charge float64) *Arrow {
	var shooter = *s
	var arrow = &Arrow{
		shooter: s,
		Charge:  charge,
	}
	arrow.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(-10, 0), 0, 10, 4),
		},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, 0), 6),
		},
		1,
	)
	arrow.SetColor(types.Colors.Red)
	arrow.SetAngle(shooter.GetAngle())
	arrow.SetPosition(shooter.GetPosition())
	//arrow.SetVelocity(shooter.GetVelocity())

	var force = types.NewVector(30, 0)
	force.RotateAbout(shooter.GetAngle(), types.NewZeroVector())
	force.Scale(charge)
	arrow.ApplyForce(force)

	return arrow
}

func (a *Arrow) Stop() {
	a.SetVelocity(types.NewZeroVector())
	(*a.spawner).RemoveArrowCollidable(a.GetID())
}

func (a *Arrow) GetShooter() int {
	return (*a.shooter).GetID()
}

func (a *Arrow) GetShooterUsername() string {
	return (*a.shooter).GetUsername()
}

func (a *Arrow) StopAndBreak() {
	(*a.spawner).RemoveArrow(a.GetID())
}

func (a *Arrow) GetClass() string {
	return "arrow"
}

func (a *Arrow) GetData(e *gob.Encoder) error {
	return e.Encode(a)
}
