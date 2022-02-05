package innocent

import (
	"encoding/gob"
	"fmt"
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
	"math"
)

type Innocent struct {
	entities.ID
	collider.Collider
	spawner *Spawner

	Username string

	input types.Input

	// the sword that is being held
	sword *Swingable
	bow   *Shootable
}

var friction = .5
var mass = 10.0

func NewInnocent(u string) *Innocent {
	var innocent = &Innocent{
		Username: u,
	}
	innocent.Setup()
	return innocent
}

func (i *Innocent) Setup() {
	i.SetupCollider(
		[]collider.Rectangle{
			collider.NewInertialRectangle(types.NewVector(22.5, 0), 0, 60, 30, 0, 10),
		},
		[]collider.Circle{
			collider.NewInertialCircle(types.NewVector(45, 15), 5, .5, 1),
			collider.NewInertialCircle(types.NewVector(0, 15), 5, .5, 1),
			collider.NewInertialCircle(types.NewVector(45, -15), 5, .5, 1),
			collider.NewInertialCircle(types.NewVector(0, -15), 5, .5, 1),
		},
		mass,
	)
	i.SetColor(types.Colors.Orange)
	i.SetFriction(friction)
}

func (i *Innocent) GetIdentifier() string {
	return i.Username
}

func (i *Innocent) ScaleMass(scale float64) {
	i.SetMass(scale * mass)
}

func (i *Innocent) ResetMass() {
	i.SetMass(mass)
}

func (i *Innocent) ShotBy(id int, username string) {
	fmt.Println(i.GetIdentifier(), "shot by", username)
	(*i.spawner).RemoveInnocent(i.GetID())
}

func (i *Innocent) SlainBy(id int, username string) {
	fmt.Println(i.GetIdentifier(), "slain by", username)
	(*i.spawner).RemoveInnocent(i.GetID())
}

func (i *Innocent) Tick() {
	var in = i.input

	var movementForce = types.NewVector(15, 0)
	if in.Forward {
		movementForce.Scale(1)
	} else if in.Backward {
		movementForce.Scale(-1)
	} else {
		movementForce.Scale(0)
	}
	movementForce.RotateAbout(i.GetAngle(), types.NewZeroVector())
	i.Collider.ApplyForce(movementForce)

	var speed = i.Velocity.Magnitude()
	var turningForce = types.NewVector(0, .005)
	turningForce.Scale(math.Sqrt(speed))
	fmt.Println(i.Velocity.Magnitude())
	if in.Left {
		turningForce.Scale(-1)
	} else if in.Right {
		turningForce.Scale(1)
	} else {
		turningForce.Scale(0)
	}
	turningForce.RotateAbout(i.GetAngle(), types.NewZeroVector())
	var frontAxle = types.NewVector(30, 0)
	frontAxle.Add(i.GetPosition())
	frontAxle.RotateAbout(i.GetAngle(), i.GetPosition())
	i.Collider.ApplyPositionalForce(turningForce, frontAxle)

	var spawner = *i.spawner
	if in.AttackClick && i.sword == nil { // initialize sword
		i.sword = spawner.SpawnSword(i)
		(*i.sword).Swing()
	} else if i.sword != nil {
		if (*i.sword).SwingCompleted() {
			spawner.DespawnSword((*i.sword).GetID())
			i.sword = nil
		}
	}

	if in.RangedClick && i.bow == nil {
		i.bow = spawner.SpawnBow(i)
	} else if i.bow != nil {
		var bow = *i.bow
		if in.AttackClick {
			bow.Cancel()
			spawner.DespawnSword(bow.GetID())
			i.bow = nil
		} else if in.RangedClick {
			bow.ChargeBow()
		} else {
			if bow.IsFired() {
				spawner.DespawnSword(bow.GetID())
				i.bow = nil
			} else {
				bow.Fire()
			}
		}
	}
}

func (i *Innocent) GetClass() string {
	return "innocent"
}

func (i *Innocent) GetData(e *gob.Encoder) error {
	return e.Encode(i)
}
