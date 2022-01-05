package collisions

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
)

type PlayerWall interface {
	GetCollider() *collider.Collider
	BounceBack()
}

type WallPlayer interface {
	CheckCollision(*collider.Collider) bool
}

func (m *Manager) AddPlayerWall(id int, p *PlayerWall) {
	m.PlayerWalls[id] = p
}

func (m *Manager) RemovePlayerWall(id int) {
	delete(m.PlayerWalls, id)
}

func (m *Manager) AddWallPlayer(id int, w *WallPlayer) {
	m.WallPlayers[id] = w
}

func (m *Manager) RemoveWallPlayer(id int) {
	delete(m.WallPlayers, id)
}

func (m *Manager) resolvePlayerWalls() {
	for _, w := range m.WallPlayers {
		var wall = *w
		for _, p := range m.PlayerWalls {
			var player = *p
			var playerCollider = player.GetCollider()
			if wall.CheckCollision(playerCollider) {
				player.BounceBack()
			}
		}
	}
}
