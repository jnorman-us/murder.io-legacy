package logic

import "github.com/josephnormandev/murder/common/types"

type Tickable interface {
	GetID() types.ID
	Tick()
}

func (m *Manager) AddTickable(id types.ID, t *Tickable) {
	m.Tickables[id] = t
}

func (m *Manager) RemoveTickable(id types.ID) {
	delete(m.Tickables, id)
}
