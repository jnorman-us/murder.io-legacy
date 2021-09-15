package classes

import "github.com/josephnormandev/murder/collider"

type Collidable interface {
	GetCollider() *collider.Collider
}
