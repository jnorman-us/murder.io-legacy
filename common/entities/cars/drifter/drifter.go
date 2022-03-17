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
	types.ID
	types.UserID
	types.Change
	entities.Health
	collider.Collider
	types.Input
	shotgunCoolDown types.CoolDown
	spawner         *Spawner
}

func NewDrifter() *Drifter {
	var drifter = &Drifter{
		shotgunCoolDown: types.NewCoolDown(40),
	}
	drifter.Setup()
	return drifter
}

func (d *Drifter) Setup() {
	d.Collider.SetupCollider(
		map[string]collider.Rectangle{
			"frontBody": collider.NewRectangle(types.NewVector(-15, 0), 0, 30, 25),
			"backBody":  collider.NewRectangle(types.NewVector(15, 0), 0, 30, 15),
		},
		map[string]collider.Circle{
			"w0": collider.NewCircle(types.NewVector(22.5, 12), 5),
			"w1": collider.NewCircle(types.NewVector(-22.5, 15), 5),
			"w2": collider.NewCircle(types.NewVector(-22.5, -15), 5),
			"w3": collider.NewCircle(types.NewVector(22.5, -12), 5),
		},
	)
	d.SetHealth(MaxHealth)
	d.Collider.SetColor(types.Colors.Red)
	d.Collider.SetMass(Mass)
	d.Collider.SetForwardFriction(Friction)
}

func (d *Drifter) Break() {
	var spawner = *d.spawner
	spawner.RemoveDrifter(d.ID)
}
