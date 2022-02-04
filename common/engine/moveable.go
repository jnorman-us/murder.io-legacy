package engine

import "github.com/josephnormandev/murder/common/collisions/collider"

type Moveable interface {
	GetID() int
	ClearBuffers()
	GetCollider() *collider.Collider
	CalculateFrictionForces(float64)
	UpdatePosition(float64)
}

func (e *Engine) AddMoveable(id int, m *Moveable) {
	e.Moveables[id] = m
}

func (e *Engine) RemoveMoveable(id int) {
	delete(e.Moveables, id)
}
