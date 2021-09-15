package collisions

import (
	"github.com/josephnormandev/murder/classes"
)

// ActorActorCollidable is a single interface that is used
// by both A and B to handle collisions
type ActorActorCollidable interface {
	classes.Identifiable
	classes.Collidable
	ActorCollidedWithActor(*ActorActorCollidable)
}
