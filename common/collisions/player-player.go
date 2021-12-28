package collisions

import (
	"fmt"
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

type PlayerPlayerCollidable interface {
	GetCollider() *collider.Collider
	GetVelocity() types.Vector
	CheckCollision(*collider.Collider) bool
	ApplyForce(types.Vector)
	GetMass() float64
	BounceBack()
}

func (m *Manager) AddPlayerPlayerCollidable(id int, c *PlayerPlayerCollidable) {
	m.PlayerPlayerCollidables[id] = c
}

func (m *Manager) RemovePlayerPlayerCollidable(id int) {
	delete(m.PlayerPlayerCollidables, id)
}

func (m *Manager) resolvePlayerPlayerCollidables() {
	for a, playerA := range m.PlayerPlayerCollidables {
		for b, playerB := range m.PlayerPlayerCollidables {
			var colliderB = (*playerB).GetCollider()
			if a != b && (*playerA).CheckCollision(colliderB) {
				var mass = (*playerA).GetMass()
				// calculate the momentum function of both bodies
				// and determine how much force to push back on both
				// of the bodies
				// mass is equal, so cancel out in m1*v1+m2*v2=...
				var combinedVelocity = (*playerA).GetVelocity()
				combinedVelocity.Add((*playerB).GetVelocity())
				combinedVelocity.Scale(0.5)

				var forceA = combinedVelocity
				var forceB = combinedVelocity
				forceA.Scale(-mass)
				forceB.Scale(mass)

				fmt.Println(forceB)

				// bounce playerA back
				//(*playerA).BounceBack()
				//(*playerB).BounceBack()
				(*playerA).ApplyForce(forceA)
				(*playerB).ApplyForce(forceB)

			}
		}
	}
}
