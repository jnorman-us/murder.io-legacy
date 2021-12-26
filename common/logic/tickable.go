package logic

type Tickable interface {
	GetID() int
	Tick()
}

func (m *Manager) AddTickable(id int, t *Tickable) {
	m.Tickables[id] = t
}

func (m *Manager) RemoveTickable(id int) {
	delete(m.Tickables, id)
}
