package collisions

type Manager struct {
	Collidables map[int32]*Collidable
}

func NewManager() *Manager {
	return &Manager{
		Collidables: map[int32]*Collidable{},
	}
}

func (m *Manager) Resolve() {
	m.ResolveCollidables()
}
