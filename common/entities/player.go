package entities

import (
	"github.com/josephnormandev/murder/common/classes"
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/events"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/world"
)

type Player struct {
	Entity
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

func (p *Player) Add() {
	p.world = world.Singleton
	p.events = events.Singleton

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
	p.events = nil
}

func (p *Player) AddInputListener() {
	var playerInputListener = events.PlayerInputListener(p)
	p.events.RegisterPlayerInputListener(&playerInputListener)
}

func (p *Player) Tick() {

}

func (p *Player) HandlePlayerInput(e events.PlayerInputEvent) {

}
