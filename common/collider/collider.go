package collider

import (
	"fmt"
	"github.com/Tarliton/collision2d"
	"github.com/josephnormandev/murder/common/types"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"image/color"
	"math"
)

type Collider struct {
	types.Change

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

	c.MarkDirty()
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
				fmt.Println("Polygon Circle")
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
				fmt.Println("Polygon Circle")
				collision.SetColliding(true)
			}
		}
		for _, b := range o.rectangles {
			var polygonB = b.getPolygon()
			var collided, _ = collision2d.TestPolygonPolygon(polygonA, polygonB)
			if collided {
				fmt.Println("Polygon P{olygon")
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
	if !newPosition.Equals(c.Position) {
		c.MarkDirty()
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
		c.MarkDirty()
	}
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

func (c *Collider) DrawHitbox(g *draw2dimg.GraphicContext) {
	for _, c := range c.circles {
		c.drawHitbox(g)
	}
	for _, r := range c.rectangles {
		r.drawHitbox(g)
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
