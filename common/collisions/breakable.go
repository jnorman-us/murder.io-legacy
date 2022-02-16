package collisions

import (
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/types"
)

// Breakable is an object that breaks upon colliding with the other object
// and imparts some damage to the object, if possible
type Breakable interface {
	Break()
	GetDamage() int
	Damage(int)
	Dead() bool
	CheckCollision(*collider.Collider) collider.Collision
}

func (m *Manager) AddBreakable(id types.ID, b *Breakable) {
	m.Breakables[id] = b
}

func (m *Manager) RemoveBreakable(id types.ID) {
	delete(m.Breakables, id)
}
