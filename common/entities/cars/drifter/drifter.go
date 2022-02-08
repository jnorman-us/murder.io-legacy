package drifter

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

var mass = 100.0
var friction = 0.5

type Drifter struct {
	types.ID
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
			collider.NewRectangle(types.NewVector(22.5, 0), 0, 60, 30),
		},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(45, 15), 5),
			collider.NewCircle(types.NewVector(0, 15), 5),
			collider.NewCircle(types.NewVector(45, -15), 5),
			collider.NewCircle(types.NewVector(0, -15), 5),
		},
	)
	d.Collider.SetMass(mass)
	d.Collider.SetFriction(friction)
}
