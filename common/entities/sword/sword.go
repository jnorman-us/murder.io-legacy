package sword

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
	"math"
)

type Sword struct {
	entities.ID
	collider.Collider

	wielder *Beholder
}

func NewSword(w *Beholder) *Sword {
	var sword = &Sword{
		wielder: w,
	}
	sword.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(30, 0), math.Pi/-4, 36, 2),
		},
		[]collider.Circle{},
		1,
	)
	sword.SetFriction((*w).GetFriction())
	sword.SetAngularFriction(.1)
	return sword
}

func (s *Sword) Tick() {
}

func (s *Sword) UpdatePosition(time float64) {
	var wielder = *s.wielder
	var copyPosition = wielder.GetPosition()
	var copyVelocity = wielder.GetVelocity()
	s.SetPosition(copyPosition)
	s.SetVelocity(copyVelocity)
	s.Collider.UpdatePosition(time)
}

func (s *Sword) Swing() {
	s.ApplyTorque(.18)
}

func (s *Sword) SwingCompleted() bool {
	return math.Abs(s.GetAngularVelocity()) < .01
}
