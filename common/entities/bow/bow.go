package bow

import (
	"encoding/gob"
	"fmt"
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

type Bow struct {
	entities.ID
	collider.Collider

	holder  *Holder
	spawner *Spawner

	Charge float64
	Fired  bool
}

func NewBow(h *Holder) *Bow {
	var holder = *h
	var bow = &Bow{
		holder: h,
		Charge: 10,
	}
	bow.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(17, 0), 0, 20, 5),
		},
		[]collider.Circle{},
		1,
	)
	bow.SetColor(types.Colors.Blue)
	bow.SetPosition(holder.GetPosition())
	bow.SetVelocity(holder.GetVelocity())
	bow.SetAngle(holder.GetAngle())
	bow.SetFriction((*h).GetFriction())
	return bow
}

func (b *Bow) Tick() {

}

func (b *Bow) UpdatePosition(time float64) {
	var holder = *b.holder
	var copyPosition = holder.GetPosition()
	var copyVelocity = holder.GetVelocity()
	var copyAngle = holder.GetAngle()
	b.SetAngle(copyAngle)
	b.SetPosition(copyPosition)
	b.SetVelocity(copyVelocity)
	b.Collider.UpdatePosition(time)
}

func (b *Bow) ChargeBow() {
	var holder = *b.holder
	if b.Charge < 50 {
		b.Charge++
		holder.ScaleMass((b.Charge / 50 * 5) + 1)
	}
}

func (b *Bow) Fire() {
	(*b.holder).ResetMass()
	(*b.spawner).SpawnArrow(b.holder, b.Charge/50)
	b.Fired = true
	b.Charge = 0
}

func (b *Bow) Cancel() {
	(*b.holder).ResetMass()
	fmt.Println("Cancelled Bow Shot")
}

func (b *Bow) IsFired() bool {
	return b.Fired
}

func (b *Bow) GetClass() string {
	return "bow"
}

func (b *Bow) GetData(e *gob.Encoder) {
	e.Encode(b)
}
