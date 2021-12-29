package collisions

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
)

type PlayerWallCollidable interface {
	GetCollider()
	BounceBack()
}

type WallPlayerCollidable interface {
	CheckCollision(*collider.Collider) bool
}
