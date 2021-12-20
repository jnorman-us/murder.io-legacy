package collider

import (
	"github.com/josephnormandev/murder/common/types"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"image/color"
)

type Collider struct {
	position        types.Vector
	angle           float64
	velocity        types.Vector
	angularVelocity float64
	rectangles      []Rectangle
	circles         []Circle
	color           color.RGBA
}

func (c *Collider) GetCollider() *Collider {
	return c
}

func (c *Collider) SetupCollider(rectangles []Rectangle, circles []Circle) {
	c.rectangles = rectangles
	c.circles = circles

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
	// then rect on rect collisions
	return colliding
}

func (c *Collider) UpdatePosition() {
	var newPosition = c.GetPosition()
	var newAngle = c.GetAngle()
	newPosition.Add(c.velocity)
	newAngle += c.angularVelocity

	c.SetPosition(newPosition)
	c.SetAngle(newAngle)
}

func (c *Collider) GetPosition() types.Vector {
	return c.position
}

func (c *Collider) SetPosition(p types.Vector) {
	c.position = p
}

func (c *Collider) GetAngle() float64 {
	return c.angle
}

func (c *Collider) SetAngle(a float64) {
	c.angle = a
}
func (c *Collider) GetVelocity() types.Vector {
	return c.velocity
}

func (c *Collider) SetVelocity(velocity types.Vector) {
	c.velocity = velocity
}

func (c *Collider) GetAngularVelocity() float64 {
	return c.angularVelocity
}

func (c *Collider) SetAngularVelocity(angularVelocity float64) {
	c.angularVelocity = angularVelocity
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

	// draw centerpoint for reference
	g.SetFillColor(color.RGBA{A: 0xff})
	g.SetStrokeColor(color.RGBA{A: 0xff})
	g.BeginPath()
	draw2dkit.Circle(g, c.position.X, c.position.Y, 2)
	g.FillStroke()
}
