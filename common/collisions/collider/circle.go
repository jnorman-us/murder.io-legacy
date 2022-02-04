package collider

import (
	"github.com/josephnormandev/murder/common/types"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"math"
)

type Circle struct {
	localPosition types.Vector
	radius        float64
	collider      *Collider

	inertial bool
	friction float64
	mass     float64
}

// NewCircle defines a circle at
// p-position,
// r-radius,
// f-friction,
// m-mass
func NewCircle(p types.Vector, r float64) Circle {
	return Circle{
		localPosition: p,
		radius:        r,
		inertial:      false,
	}
}

func NewInertialCircle(p types.Vector, r, f, m float64) Circle {
	return Circle{
		localPosition: p,
		radius:        r,

		inertial: true,
		friction: f,
		mass:     m,
	}
}

func (c *Circle) setCollider(co *Collider) {
	c.collider = co
}

func (c *Circle) checkCircleCollision(o *Circle) bool {
	var selfPos = c.getOffsetPosition()
	var otherPos = o.getOffsetPosition()
	var dist = selfPos.Distance(otherPos)

	var r1 = c.radius
	var r2 = o.radius

	var colliding = math.Pow(dist, 2) < (r1+r2)*(r1+r2)
	return colliding
}

func (c *Circle) checkRectangleCollision(r *Rectangle) bool {
	var rectPos = r.getOffsetPosition()
	var rectAngle = r.getOffsetAngle()
	var rotCirclePos = c.getOffsetPosition()
	rotCirclePos.RotateAbout(-rectAngle, rectPos)

	var radX = r.width / 2
	var radY = r.height / 2

	var closestPos = types.NewZeroVector()

	if rotCirclePos.X < rectPos.X-radX {
		closestPos.X = rectPos.X - radX
	} else if rotCirclePos.X > rectPos.X+radX {
		closestPos.X = rectPos.X + radX
	} else {
		closestPos.X = rotCirclePos.X
	}
	if rotCirclePos.Y < rectPos.Y-radY {
		closestPos.Y = rectPos.Y - radY
	} else if rotCirclePos.Y > rectPos.Y+radY {
		closestPos.Y = rectPos.Y + radY
	} else {
		closestPos.Y = rotCirclePos.Y
	}

	var distance = rotCirclePos.Distance(closestPos)
	if distance < c.radius {
		return true
	}
	return false
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

func (c *Circle) drawHitbox(g *draw2dimg.GraphicContext) {
	var position = c.getOffsetPosition()
	g.SetFillColor(c.collider.color)
	g.SetStrokeColor(c.collider.color)
	g.BeginPath()
	draw2dkit.Circle(g, position.X, position.Y, c.radius)
	g.FillStroke()
}
