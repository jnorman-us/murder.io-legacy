package engine

import (
	"github.com/josephnormandev/murder/common/types"
)

type Moveable interface {
	ClearBuffers()
	UpdatePosition(float64)
	GetPosition() types.Vector
	GetAngle() float64
}

func (e *Engine) AddMoveable(id types.ID, m *Moveable) {
	e.Moveables[id] = m
}

func (e *Engine) RemoveMoveable(id types.ID) {
	delete(e.Moveables, id)
}
