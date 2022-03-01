package collider

import (
	"github.com/Tarliton/collision2d"
	"github.com/josephnormandev/murder/common/types"
)

type Circle struct {
	localPosition types.Vector
	radius        float64
	collider      *Collider

	calculatedCircle collision2d.Circle
}

func NewCircle(p types.Vector, r float64) Circle {
	return Circle{
		localPosition: p,
		radius:        r,
	}
}

func (c *Circle) setCollider(co *Collider) {
	c.collider = co
}

func (c *Circle) getOffsetPosition() types.Vector {
	var offsetPosition = c.localPosition.Copy()
	offsetPosition.Add(c.collider.GetPosition())
	offsetPosition.RotateAbout(c.getOffsetAngle(), c.collider.GetPosition())
	return offsetPosition
}

func (c *Circle) getOffsetAngle() float64 {
	return c.collider.GetAngle()
}

func (c *Circle) calculate() {
	var offsetPosition = collision2d.Vector(c.getOffsetPosition())
	var radius = c.radius
	c.calculatedCircle = collision2d.NewCircle(offsetPosition, radius)
}

func (c *Circle) getCircle() collision2d.Circle {
	return c.calculatedCircle
}

func (c *Circle) drawHitbox() {
}
