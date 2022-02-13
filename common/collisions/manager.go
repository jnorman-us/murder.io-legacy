package collisions

import "github.com/josephnormandev/murder/common/types"

type Manager struct {
	StaticDynamics map[types.ID]*StaticDynamic
	DynamicStatics map[types.ID]*DynamicStatic
}

func NewManager() *Manager {
	return &Manager{
		StaticDynamics: map[types.ID]*StaticDynamic{},
		DynamicStatics: map[types.ID]*DynamicStatic{},
	}
}

func (m *Manager) ResolveCollisions() {
	m.resolveStaticDynamics()
}
