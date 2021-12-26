package logic

type Manager struct {
	Tickables map[int]*Tickable
}

func NewManager() *Manager {
	var manager = &Manager{
		Tickables: map[int]*Tickable{},
	}
	return manager
}

func (m *Manager) Tick() {
	for id := range m.Tickables {
		(*m.Tickables[id]).Tick()
	}
}
