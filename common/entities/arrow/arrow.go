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
	var arrow = &Arrow{
		shooter: s,
		Charge:  charge,
	}
	arrow.Setup()
	return arrow
}

func (a *Arrow) Setup() {
	a.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(-10, 0), 0, 10, 4),
		},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, 0), 6),
		},
		1,
	)
	a.SetColor(types.Colors.Red)

	if a.shooter != nil {
		var shooter = *a.shooter
		a.SetAngle(shooter.GetAngle())
		a.SetPosition(shooter.GetPosition())
		//arrow.SetVelocity(shooter.GetVelocity())

		var force = types.NewVector(30, 0)
		force.RotateAbout(shooter.GetAngle(), types.NewZeroVector())
		force.Scale(a.Charge)
		a.ApplyForce(force)
	}
}

func (a *Arrow) Stop() {
	a.SetVelocity(types.NewZeroVector())
	(*a.spawner).RemoveArrowCollidable(a.GetID())
}

func (a *Arrow) GetShooter() int {
	return (*a.shooter).GetID()
}

func (a *Arrow) GetShooterUsername() string {
	return (*a.shooter).GetIdentifier()
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
