package engine

import (
	collider2 "github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Moveable interface {
	GetID() types.ID
	CleanDirt()
	Dirty() bool
	ClearBuffers()
	GetCollider() *collider2.Collider
	UpdatePosition(float64)
}

func (e *Engine) AddMoveable(id types.ID, m *Moveable) {
	e.Moveables[id] = m
}

func (e *Engine) RemoveMoveable(id types.ID) {
	delete(e.Moveables, id)
}
