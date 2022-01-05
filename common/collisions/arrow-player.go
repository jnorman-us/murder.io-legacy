package collisions

import "github.com/josephnormandev/murder/common/collisions/collider"

type ArrowPlayer interface {
	GetCollider() *collider.Collider
	GetShooter() int
	GetShooterUsername() string
	StopAndBreak()
}

type PlayerArrow interface {
	CheckCollision(*collider.Collider) bool
	GetID() int
	ShotBy(int, string)
}

func (m *Manager) AddArrowPlayer(id int, a *ArrowPlayer) {
	m.ArrowPlayers[id] = a
}

func (m *Manager) RemoveArrowPlayer(id int) {
	delete(m.ArrowPlayers, id)
}

func (m *Manager) AddPlayerArrow(id int, p *PlayerArrow) {
	m.PlayerArrows[id] = p
}

func (m *Manager) RemovePlayerArrow(id int) {
	delete(m.PlayerArrows, id)
}

func (m *Manager) resolveArrowPlayers() {
	for _, player := range m.PlayerArrows {
		for _, arrow := range m.ArrowPlayers {
			var arrowCollider = (*arrow).GetCollider()
			if (*player).CheckCollision(arrowCollider) {
				var shooterID = (*arrow).GetShooter()
				if (*player).GetID() != shooterID {
					(*arrow).StopAndBreak()
					(*player).ShotBy(shooterID, (*arrow).GetShooterUsername())
				}
			}
		}
	}
}
