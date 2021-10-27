package collisions

import (
	"github.com/josephnormandev/murder/common/collider"
	"image/color"
)

type Collidable interface {
	GetCollider() *collider.Collider
}

func (m *Manager) AddCollidable(id int32, c *Collidable) {
	m.Collidables[id] = c
}

func (m *Manager) RemoveCollidable(id int32) {
	delete(m.Collidables, id)
}

func (m *Manager) ResolveCollidables() {
	var green = color.RGBA{
		G: 0xff,
		A: 0xff,
	}
	var red = color.RGBA{
		R: 0xff,
		A: 0xff,
	}

	for _, collidable := range m.Collidables {
		var collider = (*collidable).GetCollider()
		collider.SetColor(green)
	}
	for _, collidableA := range m.Collidables {
		var a = (*collidableA).GetCollider()
		for _, collidableB := range m.Collidables {
			var b = (*collidableB).GetCollider()
			if a != b && a.CheckCollision(b) {
				a.SetColor(red)
				b.SetColor(red)
			}
		}
	}
}
