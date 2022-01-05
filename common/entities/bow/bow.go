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

	holder  *Holder
	spawner *Spawner

	charge float64
	fired  bool
}

func NewBow(h *Holder) *Bow {
	var holder = *h
	var bow = &Bow{
		holder: h,
		charge: 10,
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

func (b *Bow) Charge() {
	var holder = *b.holder
	if b.charge < 50 {
		b.charge++
		holder.ScaleMass((b.charge / 50 * 5) + 1)
	}
}

func (b *Bow) Fire() {
	(*b.holder).ResetMass()
	(*b.spawner).SpawnArrow(b.holder, b.charge/50)
	b.fired = true
	b.charge = 0
}

func (b *Bow) Cancel() {
	(*b.holder).ResetMass()
	fmt.Println("Cancelled Bow Shot")
}

func (b *Bow) Fired() bool {
	return b.fired
}
