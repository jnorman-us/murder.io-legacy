package collider

import "github.com/josephnormandev/murder/common/types"

func (c *Collider) ClearBuffers() {
	c.forceBuffer = types.NewZeroVector()
	c.torqueBuffer = 0
}

func (c *Collider) ApplyForce(force types.Vector) {
	c.forceBuffer.Add(force)
}

func (c *Collider) ApplyPositionalForce(force types.Vector, position types.Vector) {
	c.forceBuffer.Add(force)
	var offset = c.Position.Offset(position)
	c.ApplyTorque(offset.X*force.Y - offset.Y*force.X)
}

func (c *Collider) ApplyPositionalForceAround(force types.Vector, position types.Vector, pivot types.Vector) {
	c.forceBuffer.Add(force)
	var offset = pivot.Offset(position)
	c.ApplyTorque(offset.X*force.Y - offset.Y*force.X)
}

func (c *Collider) ApplyTorque(torque float64) {
	c.torqueBuffer += torque
}
