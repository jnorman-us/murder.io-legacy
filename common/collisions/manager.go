package collisions

import "github.com/josephnormandev/murder/common/types"

type Manager struct {
	Statics    map[types.ID]*Static
	Breakables map[types.ID]*Breakable
	Dynamics   map[types.ID]*Dynamic
}

func NewManager() *Manager {
	return &Manager{
		Statics:    map[types.ID]*Static{},
		Breakables: map[types.ID]*Breakable{},
		Dynamics:   map[types.ID]*Dynamic{},
	}
}

func (m *Manager) ResolveCollisions() {
}
