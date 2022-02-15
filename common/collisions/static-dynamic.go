package collisions

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

type StaticDynamic interface {
	CheckCollision(*collider.Collider) collider.Collision
	GetPosition() types.Vector
}

type DynamicStatic interface {
	GetCollider() *collider.Collider
	ApplyForce(types.Vector)
	ApplyPositionalForce(types.Vector, types.Vector)
	GetMass() float64
	GetVelocity() types.Vector
}

func (m *Manager) resolveStaticDynamics() {
	for _, s := range m.StaticDynamics {
		var static = *s
		for _, d := range m.DynamicStatics {
			var dynamic = *d
			var collision = static.CheckCollision(dynamic.GetCollider())

			if collision.Colliding() {
				var mass = dynamic.GetMass()
				var force = dynamic.GetVelocity()
				force.Scale(-1 * mass)
				force.Scale(1)

				// var position = static.GetPosition()

				dynamic.ApplyForce(force)
			}
		}
	}
}

func (m *Manager) AddStaticDynamic(id types.ID, s *StaticDynamic) {
	m.StaticDynamics[id] = s
}

func (m *Manager) RemoveStaticDynamic(id types.ID) {
	delete(m.StaticDynamics, id)
}

func (m *Manager) AddDynamicStatic(id types.ID, d *DynamicStatic) {
	m.DynamicStatics[id] = d
}

func (m *Manager) RemoveDynamicStatic(id types.ID) {
	delete(m.DynamicStatics, id)
}
