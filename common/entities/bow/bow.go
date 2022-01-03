package bow

import (
	"fmt"
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

type Bow struct {
	entities.ID
	collider.Collider

	shooter *Shooter

	charge int
	fired  bool
}

func NewBow(s *Shooter) *Bow {
	var bow = &Bow{
		shooter: s,
		charge:  10,
	}
	bow.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(17, 0), 0, 20, 5),
		},
		[]collider.Circle{},
		1,
	)
	bow.SetFriction((*s).GetFriction())
	return bow
}

func (b *Bow) Tick() {

}

func (b *Bow) UpdatePosition(time float64) {
	var shooter = *b.shooter
	var copyPosition = shooter.GetPosition()
	var copyVelocity = shooter.GetVelocity()
	var copyAngle = shooter.GetAngle()
	b.SetAngle(copyAngle)
	b.SetPosition(copyPosition)
	b.SetVelocity(copyVelocity)
	b.Collider.UpdatePosition(time)
}

func (b *Bow) Charge() {
	var shooter = *b.shooter
	if b.charge < 50 {
		b.charge++
		shooter.ScaleMass((float64(b.charge) / 50 * 5) + 1)
	}
}

func (b *Bow) Fire() {
	fmt.Println("Firing with", b.charge/10, "charge")
	(*b.shooter).ResetMass()
	b.fired = true
	b.charge = 0
}

func (b *Bow) Fired() bool {
	return b.fired
}
