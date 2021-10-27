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
	colliding     bool
}

func NewCircle(p types.Vector, r float64) Circle {
	return Circle{
		localPosition: p,
		radius:        r,
		colliding:     false,
	}
}

func (c *Circle) setCollider(co *Collider) {
	c.collider = co
}

func (c *Circle) checkCircleCollision(o *Circle) bool {
	var selfPos = c.getOffsetPosition()

	var otherPos = o.getOffsetPosition()

	var x1 = selfPos.X
	var y1 = selfPos.Y
	var r1 = c.radius

	var x2 = otherPos.X
	var y2 = otherPos.Y
	var r2 = o.radius

	var dist = math.Abs((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))

	var colliding = dist < (r1+r2)*(r1+r2)
	if colliding == true {
		c.colliding = true
		o.colliding = true
	}
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
		r.colliding = true
		c.colliding = true
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

func (c *Circle) draw(g *draw2dimg.GraphicContext) {
	var position = c.getOffsetPosition()
	g.SetFillColor(c.collider.color)
	g.SetStrokeColor(c.collider.color)
	g.BeginPath()
	draw2dkit.Circle(g, position.X, position.Y, c.radius)
	g.FillStroke()
}
