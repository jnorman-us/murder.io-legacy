package engine

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Moveable interface {
	GetID() types.ID
	ClearBuffers()
	GetCollider() *collider.Collider
	CalculateFrictionForces(float64)
	UpdatePosition(float64)
}

func (e *Engine) AddMoveable(id types.ID, m *Moveable) {
	e.Moveables[id] = m
}

func (e *Engine) RemoveMoveable(id types.ID) {
	delete(e.Moveables, id)
}
