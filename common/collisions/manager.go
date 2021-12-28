package collisions

type Manager struct {
	Collidables             map[int]*Collidable
	PlayerPlayerCollidables map[int]*PlayerPlayerCollidable
}

func NewManager() *Manager {
	return &Manager{
		Collidables:             map[int]*Collidable{},
		PlayerPlayerCollidables: map[int]*PlayerPlayerCollidable{},
	}
}

func (m *Manager) ResolveCollisions() {
	m.resolveCollidables()
	m.resolvePlayerPlayerCollidables()
}
