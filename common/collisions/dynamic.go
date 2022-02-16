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

func (m *Manager) resolveDynamicOnStatics() {
	for _, s := range m.Statics {
		var static = *s
		for _, d := range m.Dynamics {
			var dynamic = *d
			var collision = static.CheckCollision(dynamic.GetCollider())

			if collision.Colliding() {
				var velocity = dynamic.GetVelocity()
				var restorativeForce = velocity
				restorativeForce.Scale(-1 * dynamic.GetMass())

				dynamic.ApplyForce(restorativeForce)

				// then attempt to backtrack the dynamic so it snaps right outside the static
				// preventing the drive through bug
			}
		}
	}
}
