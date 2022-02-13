package collisions

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

type BulletShootable interface {
	GetDamage() int
	GetShooter() types.ID
	GetCollider() *collider.Collider
}

type ShootableBullet interface {
	Kill()
	SetHealth(int)
	GetHealth() int
	CheckCollision() collider.Collision
}
