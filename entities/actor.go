package entities

import (
	"github.com/josephnormandev/murder/classes"
	"github.com/josephnormandev/murder/collider"
	"github.com/josephnormandev/murder/collisions"
	"github.com/josephnormandev/murder/types"
	"github.com/josephnormandev/murder/world"
)

type Actor struct {
	Entity
	displayName string
}

func NewActor(displayName string) *Actor {
	var actor = &Actor{
		displayName: displayName,
	}

	actor.Setup(
		[]collider.Rectangle{},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, 0), 2),
		},
	)
	return actor
}

func (a *Actor) AddTo(w *world.World) {
	var id = w.NextAvailableID()
	a.SetID(id)
	a.world = w

	var drawable = classes.Drawable(a)
	var moveable = classes.Moveable(a)
	var actorActor = collisions.ActorActorCollidable(a)
	var actorWall = collisions.ActorWallCollidable(a)

	w.AddDrawable(id, &drawable)
	w.AddMoveable(id, &moveable)
	w.CollisionsManager.AddActorActor(id, &actorActor)
	w.CollisionsManager.AddActorWall(id, &actorWall)
}

func (a *Actor) RemoveFrom() {
	var w = a.world
	var id = a.GetID()

	w.RemoveDrawable(id)
	w.RemoveMoveable(id)
	w.CollisionsManager.RemoveActorActor(id)
	w.CollisionsManager.RemoveActorWall(id)
}

func (a *Actor) Tick() {
}

func (a *Actor) ActorCollidedWithActor(o *collisions.ActorActorCollidable) {
	a.SetBGColor(types.White)
}

func (a *Actor) ActorCollidedWithWall(w *collisions.WallActorCollidable) {

}
