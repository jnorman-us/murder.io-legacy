package collisions

type Manager struct {
	ZombiePlayerCollidables map[int32]*ZombiePlayerCollidable
	PlayerZombieCollidables map[int32]*PlayerZombieCollidable
}

func NewManager() *Manager {
	return &Manager{
		ZombiePlayerCollidables: map[int32]*ZombiePlayerCollidable{},
		PlayerZombieCollidables: map[int32]*PlayerZombieCollidable{},
	}
}

func (m *Manager) Resolve() {
	m.ResolveZombiePlayer()
}