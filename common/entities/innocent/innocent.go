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

var angularFriction = .1
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
		[]collider.Rectangle{},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, 0), 10),
		},
		mass,
	)
	i.SetColor(types.Colors.Orange)
	i.SetAngularFriction(angularFriction)
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
	i.SetAngle(in.Direction)

	var angle = 0.0
	var movementForce = types.NewVector(15, 0)
	if in.Left && in.Forward {
		angle = math.Pi / 4 * 5
	} else if in.Left && in.Backward {
		angle = math.Pi / 4 * 3
	} else if in.Right && in.Forward {
		angle = math.Pi / 4 * 7
	} else if in.Right && in.Backward {
		angle = math.Pi / 4 * 1
	} else if in.Left {
		angle = math.Pi
	} else if in.Forward {
		angle = math.Pi / 2 * 3
	} else if in.Backward {
		angle = math.Pi / 2
	} else if in.Right {
		angle = 0
	} else {
		movementForce.Scale(0)
	}

	movementForce.RotateAbout(angle, types.NewZeroVector())
	i.Collider.ApplyForce(movementForce)

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
