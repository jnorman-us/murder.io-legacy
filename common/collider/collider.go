package collider

import (
	"github.com/josephnormandev/murder/common/types"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"image/color"
	"math"
)

type Collider struct {
	types.Change
	mass            float64
	forceBuffer     types.Vector
	torqueBuffer    float64
	Position        types.Vector
	Angle           float64
	Velocity        types.Vector
	AngularVelocity float64
	friction        float64
	rectangles      []*Rectangle
	circles         []*Circle
	color           color.RGBA
}

func (c *Collider) GetCollider() *Collider {
	return c
}

func (c *Collider) SetupCollider(rectangles []Rectangle, circles []Circle) {
	c.rectangles = []*Rectangle{}
	c.circles = []*Circle{}

	for i := range rectangles {
		var rectangle = &rectangles[i]
		rectangle.setCollider(c)
		c.rectangles = append(c.rectangles, rectangle)
	}

	for i := range circles {
		var circle = &circles[i]
		circle.setCollider(c)
		c.circles = append(c.circles, circle)
	}

	c.SetColor(color.RGBA{
		G: 0xff,
		A: 0xff,
	})
	c.Set()
}

func (c *Collider) CheckCollision(o *Collider) Collision {
	var collision = NewCollision()

	// circle on circle collisions
	for _, c := range c.circles {
		for _, o := range o.circles {
			var circle = *c
			var otherCircle = o
			if circle.checkCircleCollision(otherCircle) {
				collision.SetColliding(true)
			}
		}
	}
	// then circle on rect collisions
	for _, r := range c.rectangles {
		for _, c := range o.circles {
			var rectangle = *r
			var otherCircle = c
			if rectangle.checkCircleCollision(otherCircle) {
				collision.SetColliding(true)
			}
		}
	}

	for _, c := range c.circles {
		for _, r := range o.rectangles {
			var circle = *c
			var otherRectangle = r
			if circle.checkRectangleCollision(otherRectangle) {
				collision.SetColliding(true)
			}
		}
	}

	/*
		// then rect on rect collisions
		for i := range c.rectangles {
			for j := range o.rectangles {
				var rectangle = &c.rectangles[i]
				var otherRectangle = &o.rectangles[j]
				if rectangle.checkRectangleCollision(otherRectangle) {
					colliding = true
				}
			}
		}
	*/

	return collision
}

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

func (c *Collider) CalculateFrictionForces(time float64) {

}

func (c *Collider) UpdatePosition(time float64) {
	var timeSquared = math.Pow(time*1, 2)
	var frictionAir = 1 - c.friction

	var acceleration = c.forceBuffer
	acceleration.Scale(1 / c.mass)
	acceleration.Scale(timeSquared)

	var newVelocity = c.Velocity
	newVelocity.Scale(frictionAir)
	newVelocity.Add(acceleration)
	c.SetVelocity(newVelocity)

	var newPosition = c.Position
	newPosition.Add(newVelocity)
	if !newPosition.Equals(c.Position) {
		c.Set()
	}
	c.SetPosition(newPosition)

	var angularAcceleration = c.torqueBuffer
	angularAcceleration /= c.mass
	angularAcceleration *= timeSquared

	var newAngularVelocity = c.AngularVelocity
	newAngularVelocity *= frictionAir
	newAngularVelocity += angularAcceleration
	c.SetAngularVelocity(newAngularVelocity)

	var newAngle = c.Angle
	newAngle += newAngularVelocity
	if newAngle != c.Angle {
		c.Set()
	}
	c.SetAngle(newAngle)
}

func (c *Collider) BounceBack() {
	var newPosition = c.GetPosition()
	var newVelocity = c.GetVelocity()
	newVelocity.Scale(-1 / (1 - c.friction))
	newPosition.Add(newVelocity)

	c.SetPosition(newPosition)
}

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
}

func (c *Collider) GetAngle() float64 {
	return c.Angle
}

func (c *Collider) SetAngle(a float64) {
	c.Angle = a
}
func (c *Collider) GetVelocity() types.Vector {
	return c.Velocity
}

func (c *Collider) SetVelocity(velocity types.Vector) {
	c.Velocity = velocity
}

func (c *Collider) GetFriction() float64 {
	return c.friction
}

func (c *Collider) SetFriction(coefficient float64) {
	c.friction = coefficient
}

func (c *Collider) GetAngularVelocity() float64 {
	return c.AngularVelocity
}

func (c *Collider) SetAngularVelocity(angularVelocity float64) {
	c.AngularVelocity = angularVelocity
}

func (c *Collider) SetColor(co color.RGBA) {
	c.color = co
}

func (c *Collider) DrawHitbox(g *draw2dimg.GraphicContext) {
	for _, c := range c.circles {
		(*c).drawHitbox(g)
	}
	for _, r := range c.rectangles {
		(*r).drawHitbox(g)
	}

	var directionPoint = c.Position
	directionPoint.Add(types.NewVector(20, 0))
	directionPoint.RotateAbout(c.Angle, c.Position)

	// draw centerpoint for reference
	g.SetFillColor(color.RGBA{A: 0xff})
	g.SetStrokeColor(color.RGBA{A: 0xff})
	g.BeginPath()
	draw2dkit.Circle(g, c.Position.X, c.Position.Y, 2)
	draw2dkit.Circle(g, directionPoint.X, directionPoint.Y, 0)

	g.FillStroke()
}

func (c *Collider) CopyKinetics(o Collider) {
	c.Position = o.Position
	c.Angle = o.Angle
	c.Velocity = o.Velocity
	c.AngularVelocity = o.AngularVelocity
	c.friction = o.friction
}
