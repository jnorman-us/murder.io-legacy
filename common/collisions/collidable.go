package collisions

import (
	"github.com/josephnormandev/murder/common/collider"
)

type Collidable interface {
	GetCollider() *collider.Collider
}
