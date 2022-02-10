package engine

import "github.com/josephnormandev/murder/common/types"

type Moveable interface {
	GetID() types.ID

	SetPosition(types.Vector)
	SetAngle(float64)

	GetPosition() types.Vector
	GetAngle() float64

	UpdatePosition(float64)
}

func (m *Manager) AddMoveable(id types.ID, mo *Moveable) {
	m.moveables[id] = mo
}

func (m *Manager) RemoveMoveable(id types.ID) {
	delete(m.moveables, id)
}
