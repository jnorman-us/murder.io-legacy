package sword

import (
	"encoding/gob"
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
	sword.Setup()
	return sword
}

func (s *Sword) Setup() {
	s.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(30, 0), math.Pi/-4, 30, 2),
		},
		[]collider.Circle{},
		1,
	)
	s.SetColor(types.Colors.Blue)
	s.SetFriction(.2)

	if s.wielder != nil {
		var wielder = *s.wielder
		s.SetPosition(wielder.GetPosition())
		s.SetVelocity(wielder.GetVelocity())
		s.SetAngle(wielder.GetAngle())
		s.SetFriction(wielder.GetFriction())
	}
}

func (s *Sword) Tick() {
}

func (s *Sword) UpdatePosition(time float64) {
	if s.wielder != nil {
		var wielder = *s.wielder
		var copyPosition = wielder.GetPosition()
		var copyVelocity = wielder.GetVelocity()
		s.SetPosition(copyPosition)
		s.SetVelocity(copyVelocity)
		s.Collider.UpdatePosition(time)
	}
}

func (s *Sword) Swing() {
	s.ApplyTorque(.3)
}

func (s *Sword) SwingCompleted() bool {
	return math.Abs(s.GetAngularVelocity()) < .01
}

func (s *Sword) GetWielder() int {
	return (*s.wielder).GetID()
}

func (s *Sword) GetWielderUsername() string {
	return (*s.wielder).GetIdentifier()
}

func (s *Sword) GetClass() string {
	return "sword"
}

func (s *Sword) GetData(e *gob.Encoder) error {
	return e.Encode(s)
}
