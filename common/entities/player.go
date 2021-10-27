package entities

import (
	"github.com/josephnormandev/murder/common/classes"
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/collisions"
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

func (p *Player) AddTo(w *world.World) {
	var id = w.NextAvailableID()

	var identifiable = classes.Identifiable(p)
	var moveable = classes.Moveable(p)
	var collidable = collisions.Collidable(p)

	w.AddIdentifiable(id, &identifiable)
	w.AddMoveable(id, &moveable)
	w.CollisionsManager.AddCollidable(id, &collidable)
}

func (p *Player) RemoveFrom() {
	var id = p.GetID()
	var w = p.world

	w.CollisionsManager.RemoveCollidable(id)
	w.RemoveIdentifiable(id)
	w.RemoveMoveable(id)

	w = nil
}

func (p *Player) Tick() {

}
