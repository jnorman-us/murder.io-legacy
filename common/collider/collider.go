package collider

import (
	"github.com/Tarliton/collision2d"
	"github.com/josephnormandev/murder/common/types"
	"image/color"
	"math"
)

type Collider struct {
	mass         float64
	forceBuffer  types.Vector
	torqueBuffer float64

	Position        types.Vector
	Angle           float64
	Velocity        types.Vector
	AngularVelocity float64

	forwardFriction float64
	lateralFriction float64
	angularFriction float64

	rectangles map[string]*Rectangle
	circles    map[string]*Circle
	color      color.RGBA
}

func (c *Collider) GetCollider() *Collider {
	return c
}

func (c *Collider) SetupCollider(rectangles map[string]Rectangle, circles map[string]Circle) {
	c.rectangles = map[string]*Rectangle{}
	c.circles = map[string]*Circle{}

	for id, r := range rectangles {
		var rectangle = r
		rectangle.setCollider(c)
		c.rectangles[id] = &rectangle
	}

	for id, ci := range circles {
		var circle = ci
		circle.setCollider(c)
		c.circles[id] = &circle
	}
}

func (c *Collider) CheckCollision(o *Collider) Collision {
	var collision = Collision{}
	for _, a := range c.circles {
		var circleA = a.getCircle()
		for _, b := range o.circles {
			var circleB = b.getCircle()
			var collided, _ = collision2d.TestCircleCircle(circleA, circleB)
			if collided {
				collision.SetColliding(true)
			}
		}
		for _, b := range o.rectangles {
			var polygonB = b.getPolygon()
			var collided, _ = collision2d.TestCirclePolygon(circleA, polygonB)
			if collided {
				collision.SetColliding(true)
			}
		}
	}

	for _, a := range c.rectangles {
		var polygonA = a.getPolygon()
		for _, b := range o.circles {
			var circleB = b.getCircle()
			var collided, _ = collision2d.TestPolygonCircle(polygonA, circleB)
			if collided {
				collision.SetColliding(true)
			}
		}
		for _, b := range o.rectangles {
			var polygonB = b.getPolygon()
			var collided, _ = collision2d.TestPolygonPolygon(polygonA, polygonB)
			if collided {
				collision.SetColliding(true)
			}
		}
	}
	return collision
}

func (c *Collider) UpdatePosition(time float64) {
	var timeSquared = math.Pow(time*1, 2)
	var frictionAir = 1 - c.forwardFriction

	var acceleration = c.forceBuffer
	acceleration.Scale(1 / c.mass)
	acceleration.Scale(timeSquared)

	var newVelocity = c.Velocity
	newVelocity.Scale(frictionAir)
	newVelocity.Add(acceleration)
	c.SetVelocity(newVelocity)

	var newPosition = c.Position
	newPosition.Add(newVelocity)
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
	c.SetAngle(newAngle)
	c.CalculateHitbox()
}

func (c *Collider) CalculateHitbox() {
	for _, circle := range c.circles {
		circle.calculate()
	}
	for _, rectangle := range c.rectangles {
		rectangle.calculate()
	}
}

func (c *Collider) SetColor(co color.RGBA) {
	c.color = co
}

func (c *Collider) GetColor() color.RGBA {
	return c.color
}
