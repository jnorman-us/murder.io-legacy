package collisions

import (
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Static interface {
	CheckCollision(*collider.Collider) collider.Collision
}

func (m *Manager) AddStatic(id types.ID, s *Static) {
	m.Statics[id] = s
}

func (m *Manager) RemoveStatic(id types.ID) {
	delete(m.Statics, id)
}
