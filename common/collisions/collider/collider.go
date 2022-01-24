package collider

import (
	"github.com/josephnormandev/murder/common/types"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"image/color"
	"math"
)

type Collider struct {
	mass            float64
	Position        types.Vector
	Angle           float64
	Velocity        types.Vector
	AngularVelocity float64
	Friction        float64
	AngularFriction float64
	rectangles      []Rectangle
	circles         []Circle
	color           color.RGBA
}

func (c *Collider) GetCollider() *Collider {
	return c
}

func (c *Collider) SetupCollider(rectangles []Rectangle, circles []Circle, mass float64) {
	c.rectangles = rectangles
	c.circles = circles
	c.mass = mass

	for i := range c.rectangles {
		var rectangle = &c.rectangles[i]
		rectangle.setCollider(c)
	}

	for i := range c.circles {
		var circle = &c.circles[i]
		circle.setCollider(c)
	}
	c.SetColor(color.RGBA{
		G: 0xff,
		A: 0xff,
	})
}

func (c *Collider) Clear() {
	for i := range c.circles {
		var circle = &c.circles[i]
		circle.colliding = false
	}
	for i := range c.rectangles {
		var rectangle = &c.rectangles[i]
		rectangle.colliding = false
	}
}

func (c *Collider) CheckCollision(o *Collider) bool {
	// circle on circle collisions
	var colliding = false
	for i := range c.circles {
		for j := range o.circles {
			var circle = &c.circles[i]
			var otherCircle = &o.circles[j]
			if circle.checkCircleCollision(otherCircle) {
				colliding = true
			}
		}
	}
	// then circle on rect collisions
	for i := range c.rectangles {
		for j := range o.circles {
			var rectangle = &c.rectangles[i]
			var otherCircle = &o.circles[j]
			if rectangle.checkCircleCollision(otherCircle) {
				colliding = true
			}
		}
	}

	for i := range c.circles {
		for j := range o.rectangles {
			var circle = &c.circles[i]
			var otherRectangle = &o.rectangles[j]
			if circle.checkRectangleCollision(otherRectangle) {
				colliding = true
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

	return colliding
}

func (c *Collider) ApplyForce(force types.Vector) {
	force.Scale(1 / c.mass)
	c.Velocity.Add(force)
}

func (c *Collider) ApplyTorque(torque float64) {
	c.AngularVelocity += torque / c.mass
}

func (c *Collider) UpdatePosition(tick float64) {
	var scaledVelocity = c.GetVelocity()
	var scaledAngularVelocity = c.GetAngularVelocity()
	scaledVelocity.Scale(tick)
	scaledAngularVelocity *= tick

	var scaledFriction = 1 - c.Friction*tick
	var scaledAngularFriction = 1 - c.AngularFriction*tick

	var newPosition = c.GetPosition()
	var newAngle = c.GetAngle()
	var newVelocity = c.GetVelocity()
	var newAngularVelocity = c.GetAngularVelocity()

	newPosition.Add(scaledVelocity)
	newVelocity.Scale(scaledFriction)
	newAngle += scaledAngularVelocity
	newAngularVelocity *= scaledAngularFriction

	c.SetPosition(newPosition)
	c.SetAngle(newAngle)
	c.SetVelocity(newVelocity)
	c.SetAngularVelocity(newAngularVelocity)
}

func (c *Collider) BounceBack() {
	var newPosition = c.GetPosition()
	var newVelocity = c.GetVelocity()
	newVelocity.Scale(-1 / (1 - c.Friction))
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
	return c.Friction
}

func (c *Collider) SetFriction(coefficient float64) {
	c.Friction = coefficient
}

func (c *Collider) GetAngularVelocity() float64 {
	return c.AngularVelocity
}

func (c *Collider) SetAngularVelocity(angularVelocity float64) {
	c.AngularVelocity = angularVelocity
}

func (c *Collider) SetAngularFriction(coefficient float64) {
	c.AngularFriction = coefficient
}

func (c *Collider) SetColor(co color.RGBA) {
	c.color = co
}

func (c *Collider) DrawHitbox(g *draw2dimg.GraphicContext) {
	for _, circle := range c.circles {
		circle.drawHitbox(g)
	}
	for _, rectangle := range c.rectangles {
		rectangle.drawHitbox(g)
	}

	var directionPoint = c.Position
	directionPoint.Add(types.NewVector(0, 20))
	directionPoint.RotateAbout(c.Angle-math.Pi/2, c.Position)

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
	c.Friction = o.Friction
	c.AngularFriction = o.AngularFriction
}
