package dimetrodon

import (
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

const Mass = 120.0
const Friction = 0.3
const MaxHealth = 250

type Dimetrodon struct {
	entities.Health
	types.ID
	types.UserID
	collider.Collider
	types.Input
	gattlingGunCoolDown types.CoolDown
	spawner             *Spawner
}

func NewDimetrodon() *Dimetrodon {
	var dimetrodon = &Dimetrodon{
		gattlingGunCoolDown: types.NewCoolDown(5),
	}
	dimetrodon.Setup()
	return dimetrodon
}

func (d *Dimetrodon) Setup() {
	d.Collider.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(0, 0), 0, 60, 25),
		},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(22.5, 15), 5),
			collider.NewCircle(types.NewVector(-22.5, 15), 5),
			collider.NewCircle(types.NewVector(-22.5, -15), 5),
			collider.NewCircle(types.NewVector(22.5, -15), 5),
		},
	)
	d.Health.SetHealth(MaxHealth)
	d.Collider.SetColor(types.Colors.Red)
	d.Collider.SetMass(Mass)
	d.Collider.SetFriction(Friction)
}

func (d *Dimetrodon) Break() {
	var spawner = *d.spawner
	spawner.RemoveDimetrodon(d.ID)
}
