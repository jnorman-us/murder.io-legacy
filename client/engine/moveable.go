package engine

import "github.com/josephnormandev/murder/common/types"

type Moveable interface {
	GetID() int

	SetPosition(types.Vector)
	SetAngle(float64)

	GetPosition() types.Vector
	GetAngle() float64

	UpdatePosition(float64)
}

func (m *Manager) AddMoveable(id int, mo *Moveable) {
	m.moveables[id] = mo
}

func (m *Manager) RemoveMoveable(id int) {
	delete(m.moveables, id)
}
