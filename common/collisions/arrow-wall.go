package collisions

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
)

type ArrowWall interface {
	GetCollider() *collider.Collider
	Stop()
}

type WallArrow interface {
	CheckCollision(*collider.Collider) bool
}

func (m *Manager) AddArrowWall(id int, a *ArrowWall) {
	m.ArrowWalls[id] = a
}

func (m *Manager) RemoveArrowWall(id int) {
	delete(m.ArrowWalls, id)
}

func (m *Manager) AddWallArrow(id int, w *WallArrow) {
	m.WallArrows[id] = w
}

func (m *Manager) RemoveWallArrow(id int) {
	delete(m.WallArrows, id)
}

func (m *Manager) resolveArrowWalls() {
	for _, w := range m.WallArrows {
		var wall = *w
		for _, a := range m.ArrowWalls {
			var arrow = *a
			var arrowCollider = arrow.GetCollider()
			if wall.CheckCollision(arrowCollider) {
				arrow.Stop()
			}
		}
	}
}
