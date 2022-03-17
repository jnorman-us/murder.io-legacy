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
	types.ID
	types.UserID
	types.Change
	entities.Health
	collider.Collider
	types.Input
	gattlingGunCoolDown types.CoolDown
	spawner             *Spawner
}

func NewDimetrodon() *Dimetrodon {
	var dimetrodon = &Dimetrodon{
		gattlingGunCoolDown: types.NewCoolDown(20),
	}
	dimetrodon.Setup()
	return dimetrodon
}

func (d *Dimetrodon) Setup() {
	d.Collider.SetupCollider(
		map[string]collider.Rectangle{
			"body": collider.NewRectangle(types.NewVector(0, 0), 0, 60, 25),
		},
		map[string]collider.Circle{
			"w0": collider.NewCircle(types.NewVector(22.5, 15), 5),
			"w1": collider.NewCircle(types.NewVector(-22.5, 15), 5),
			"w2": collider.NewCircle(types.NewVector(-22.5, -15), 5),
			"w3": collider.NewCircle(types.NewVector(22.5, -15), 5),
		},
	)
	d.Health.SetHealth(MaxHealth)
	d.Collider.SetColor(types.Colors.Red)
	d.Collider.SetMass(Mass)
	d.Collider.SetForwardFriction(Friction)
}

func (d *Dimetrodon) Break() {
	var spawner = *d.spawner
	spawner.RemoveDimetrodon(d.ID)
}
