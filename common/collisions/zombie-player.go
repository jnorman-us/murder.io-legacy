package collisions

type ZombiePlayerCollidable interface {
	Collidable
	HitAPlayer(*PlayerZombieCollidable)
}


type PlayerZombieCollidable interface {
	Collidable
	TakeZombieDamage(*ZombiePlayerCollidable)
}

func (m *Manager) AddZombiePlayer(id int32, z *ZombiePlayerCollidable) {
	m.ZombiePlayerCollidables[id] = z
}

func (m *Manager) AddPlayerZombie(id int32, p *PlayerZombieCollidable) {
	m.PlayerZombieCollidables[id] = p
}

func (m *Manager) RemoveZombiePlayer(id int32) {
	delete(m.ZombiePlayerCollidables, id)
}

func (m *Manager) ResolveZombiePlayer() {

}