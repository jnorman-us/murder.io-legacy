package collisions

import (
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Dynamic interface {
	GetCollider() *collider.Collider
	ApplyForce(types.Vector)
	ApplyPositionalForce(types.Vector, types.Vector)
	GetMass() float64
	GetVelocity() types.Vector
}

func (m *Manager) AddDynamic(id types.ID, d *Dynamic) {
	m.Dynamics[id] = d
}

func (m *Manager) RemoveDynamic(id types.ID) {
	delete(m.Dynamics, id)
}
