package engine

import (
	collider2 "github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"time"
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
	var offsetTime = time.Now().Sub(e.lastSendTime)
	e.Moveables[id] = m
	e.kinetics[id] = &packets.Kinetic{
		Offset: int32(offsetTime),
	}
}

func (e *Engine) RemoveMoveable(id types.ID) {
	delete(e.Moveables, id)
	delete(e.kinetics, id)
}
