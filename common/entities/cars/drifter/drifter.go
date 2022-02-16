package drifter

import (
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

const Mass = 100.0
const Friction = 0.3
const MaxHealth = 200

type Drifter struct {
	entities.Health
	types.ID
	types.UserID
	collider.Collider
	types.Input
	shotgunCoolDown types.CoolDown
	spawner         *Spawner
}

func NewDrifter() *Drifter {
	var drifter = &Drifter{
		shotgunCoolDown: types.NewCoolDown(60),
	}
	drifter.Setup()
	return drifter
}

func (d *Drifter) Setup() {
	d.Collider.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(-15, 0), 0, 30, 25),
			collider.NewRectangle(types.NewVector(15, 0), 0, 30, 15),
		},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(22.5, 12), 5),
			collider.NewCircle(types.NewVector(-22.5, 15), 5),
			collider.NewCircle(types.NewVector(-22.5, -15), 5),
			collider.NewCircle(types.NewVector(22.5, -12), 5),
		},
	)
	d.Health.SetHealth(MaxHealth)
	d.Collider.SetColor(types.Colors.Red)
	d.Collider.SetMass(Mass)
	d.Collider.SetFriction(Friction)
}

func (d *Drifter) Break() {
	var spawner = *d.spawner
	spawner.RemoveDrifter(d.ID)
}
