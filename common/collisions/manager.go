package collisions

type Manager struct {
	Collidables map[int]*Collidable
}

func NewManager() *Manager {
	return &Manager{
		Collidables: map[int]*Collidable{},
	}
}

func (m *Manager) ResolveCollisions() {
	m.resolveCollidables()
}
