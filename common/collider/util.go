package collider

import "github.com/josephnormandev/murder/common/types"

func (c *Collider) GetMass() float64 {
	return c.mass
}

func (c *Collider) SetMass(mass float64) {
	c.mass = mass
}

func (c *Collider) GetPosition() types.Vector {
	return c.Position
}

func (c *Collider) SetPosition(p types.Vector) {
	c.Position = p
	c.CalculateHitbox()
}

func (c *Collider) GetAngle() float64 {
	return c.Angle
}

func (c *Collider) SetAngle(a float64) {
	c.Angle = a
	c.CalculateHitbox()
}
func (c *Collider) GetVelocity() types.Vector {
	return c.Velocity
}

func (c *Collider) SetVelocity(velocity types.Vector) {
	c.Velocity = velocity
}

func (c *Collider) GetAngularVelocity() float64 {
	return c.AngularVelocity
}

func (c *Collider) SetAngularVelocity(angularVelocity float64) {
	c.AngularVelocity = angularVelocity
}

func (c *Collider) GetForwardFriction() float64 {
	return c.forwardFriction
}

func (c *Collider) SetForwardFriction(f float64) {
	c.forwardFriction = f
}

func (c *Collider) GetLateralFriction() float64 {
	return c.lateralFriction
}

func (c *Collider) SetLateralFriction(f float64) {
	c.lateralFriction = f
}

func (c *Collider) GetAngularFriction() float64 {
	return c.angularFriction
}

func (c *Collider) SetAngularFriction(f float64) {
	c.angularFriction = f
}
