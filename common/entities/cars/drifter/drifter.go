package drifter

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

var mass = 100.0
var friction = 0.3

type Drifter struct {
	types.ID
	types.UserID
	collider.Collider
	types.Input
}

func NewDrifter() *Drifter {
	var drifter = &Drifter{}
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
	d.Collider.SetColor(types.Colors.Red)
	d.Collider.SetMass(mass)
	d.Collider.SetFriction(friction)
}

func (d *Drifter) Tick() {

}
