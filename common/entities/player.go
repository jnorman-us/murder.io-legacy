package entities

import (
	"github.com/josephnormandev/murder/common/classes"
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/events"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
	"math"
)

type Player struct {
	Entity
	playerInputs events.PlayerInputEvent
}

func NewPlayer() *Player {
	var player = &Player{}

	player.Setup(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(0, 50), 0, 30, 30),
		},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, -50), 15),
		},
	)
	player.SetStats(20, 5)

	return player
}

func (p *Player) Add(w *world.World, e *events.Manager) {
	p.world = w
	p.eventsManager = e

	var id = p.world.NextAvailableID()

	var identifiable = classes.Identifiable(p)
	var moveable = classes.Moveable(p)
	var collidable = collisions.Collidable(p)

	p.world.AddIdentifiable(id, &identifiable)
	p.world.AddMoveable(id, &moveable)
	p.world.CollisionsManager.AddCollidable(id, &collidable)
}

func (p *Entity) Remove() {
	var id = p.GetID()

	p.world.CollisionsManager.RemoveCollidable(id)
	p.world.RemoveIdentifiable(id)
	p.world.RemoveMoveable(id)

	p.world = nil
	p.eventsManager = nil
}

func (p *Player) AddInputListener() {
	var playerInputListener = events.PlayerInputListener(p)
	p.eventsManager.RegisterPlayerInputListener(&playerInputListener)
}

func (p *Player) Tick() {
	var velocityVector types.Vector
	var aboutVector = types.NewZeroVector()
	if p.playerInputs.Dead() {
		velocityVector = types.NewZeroVector()
	} else {
		velocityVector = types.NewVector(2, 0)
	}

	if p.playerInputs.Forward && p.playerInputs.Left {
		velocityVector.RotateAbout(math.Pi/-4*3, aboutVector)
	} else if p.playerInputs.Forward && p.playerInputs.Right {
		velocityVector.RotateAbout(math.Pi/-4, aboutVector)
	} else if p.playerInputs.Backward && p.playerInputs.Right {
		velocityVector.RotateAbout(math.Pi/4, aboutVector)
	} else if p.playerInputs.Backward && p.playerInputs.Left {
		velocityVector.RotateAbout(math.Pi/4*3, aboutVector)
	} else if p.playerInputs.Forward {
		velocityVector.RotateAbout(math.Pi/-2, aboutVector)
	} else if p.playerInputs.Backward {
		velocityVector.RotateAbout(math.Pi/2, aboutVector)
	} else if p.playerInputs.Left {
		velocityVector.RotateAbout(math.Pi/-1, aboutVector)
	}
	p.SetVelocity(velocityVector)
}

func (p *Player) HandlePlayerInput(e events.PlayerInputEvent) {
	p.playerInputs = e
}
