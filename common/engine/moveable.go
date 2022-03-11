package engine

import (
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
)

type Moveable interface {
	GetID() types.ID
	ClearBuffers()
	GetCollider() *collider.Collider
	UpdatePosition(float64)
}

func (e *Engine) AddMoveable(id types.ID, m *Moveable) {
	e.Moveables[id] = m
	e.kinetics[id] = &packets.Kinetic{
		Offset: e.time.GetOffset(),
	}
}

func (e *Engine) RemoveMoveable(id types.ID) {
	delete(e.Moveables, id)
	delete(e.kinetics, id)
}
