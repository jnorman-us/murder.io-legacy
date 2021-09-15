package collisions

import "github.com/josephnormandev/murder/classes"

// ActorWallCollidable is a pair interface where
// ActorWall... is the Actor that runs into the Wall
type ActorWallCollidable interface {
	classes.Identifiable
	classes.Collidable
	ActorCollidedWithWall(*WallActorCollidable)
}

// WallActorCollidable is a pair interface where
// ActorWall... is the Wall that runs into the Actor
type WallActorCollidable interface {
	classes.Identifiable
	classes.Collidable
	WallCollidedWithActor(*ActorWallCollidable)
}
