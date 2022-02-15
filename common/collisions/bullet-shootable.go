package collisions

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

type BulletShootable interface {
	Destroy()
	GetDamage() int
	GetShooter() types.ID
	GetCollider() *collider.Collider
}

type ShootableBullet interface {
	GetID() types.ID
	Kill()
	Damage(int)
	Dead() bool
	CheckCollision(*collider.Collider) collider.Collision
}

func (m *Manager) resolveBulletShootables() {
	for _, s := range m.ShootableBullets {
		var shootable = *s
		for _, b := range m.BulletShootables {
			var bullet = *b

			if bullet.GetShooter() == shootable.GetID() {
				continue
			}

			var collision = shootable.CheckCollision(bullet.GetCollider())
			if collision.Colliding() {
				var damage = bullet.GetDamage()
				shootable.Damage(damage)

				if shootable.Dead() {
					shootable.Kill()
				}
				bullet.Destroy()
			}
		}
	}
}

func (m *Manager) AddBulletShootable(id types.ID, b *BulletShootable) {
	m.BulletShootables[id] = b
}

func (m *Manager) RemoveBulletShootable(id types.ID) {
	delete(m.BulletShootables, id)
}

func (m *Manager) AddShootableBullet(id types.ID, s *ShootableBullet) {
	m.ShootableBullets[id] = s
}

func (m *Manager) RemoveShootableBullet(id types.ID) {
	delete(m.ShootableBullets, id)
}
