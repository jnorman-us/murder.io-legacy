package collider

import (
	"fmt"
	"github.com/josephnormandev/murder/common/types"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"math"
)

type Rectangle struct {
	localPosition types.Vector
	localAngle    float64
	width         float64
	height        float64
	collider      *Collider

	inertial bool
	friction float64
	mass     float64
}

func NewRectangle(p types.Vector, a, w, h float64) Rectangle {
	return Rectangle{
		localPosition: p,
		localAngle:    a,
		width:         w,
		height:        h,
		inertial:      false,
	}
}

func NewInertialRectangle(p types.Vector, a, w, h, f, m float64) Rectangle {
	return Rectangle{
		localPosition: p,
		localAngle:    a,
		width:         w,
		height:        h,

		inertial: true,
		friction: f,
		mass:     m,
	}
}

func (r *Rectangle) setCollider(c *Collider) {
	r.collider = c
}

func (r *Rectangle) checkCircleCollision(o *Circle) bool {
	return o.checkRectangleCollision(r)
}

func (r *Rectangle) checkRectangleCollision(o *Rectangle) bool {
	if r.checkBoundingRadiusCollision(o) == false {
		return false
	}

	var radiusA = types.NewVector(r.width, r.height)
	radiusA.Scale(.5)
	var radiusB = types.NewVector(o.width, o.height)
	radiusB.Scale(.5)
	var flipX = types.NewVector(-1, 1)
	var flipY = types.NewVector(1, -1)

	// bottom right (+, +)
	var aBR = r.getOffsetPosition()
	var bBR = o.getOffsetPosition()
	aBR.Add(radiusA)
	bBR.Add(radiusB)

	radiusA.MultiplyBy(flipX)
	radiusB.MultiplyBy(flipX)

	// bottom left (-, +)
	var aBL = r.getOffsetPosition()
	var bBL = o.getOffsetPosition()
	aBL.Add(radiusA)
	bBL.Add(radiusB)

	radiusA.MultiplyBy(flipY)
	radiusB.MultiplyBy(flipY)

	var aTL = r.getOffsetPosition()
	var bTL = r.getOffsetPosition()
	aTL.Add(radiusA)
	bTL.Add(radiusB)

	radiusA.MultiplyBy(flipX)
	radiusB.MultiplyBy(flipX)

	var aTR = r.getOffsetPosition()
	var bTR = r.getOffsetPosition()
	aTR.Add(radiusA)
	bTR.Add(radiusB)

	fmt.Println(aBR, aBL, aTL, aTR)

	return false
}

func (r *Rectangle) checkBoundingRadiusCollision(o *Rectangle) bool {
	var selfPos = r.getOffsetPosition()
	var otherPos = o.getOffsetPosition()
	var dist = selfPos.Distance(otherPos)

	var r1 = r.getBoundingRadius()
	var r2 = o.getBoundingRadius()

	var colliding = math.Pow(dist, 2) < (r1+r2)*(r1+r2)
	return colliding
}

func (r *Rectangle) getBoundingRadius() float64 {
	return math.Max(r.width, r.height) / 2
}

func (r *Rectangle) getOffsetPosition() types.Vector {
	var offsetPosition = r.localPosition.Copy()
	var offsetAngle = r.getOffsetAngle()
	offsetPosition.Add(r.collider.GetPosition())
	offsetPosition.RotateAbout(offsetAngle, r.collider.GetPosition())
	return offsetPosition
}

func (r *Rectangle) getOffsetAngle() float64 {
	return r.localAngle + r.collider.GetAngle()
}

func (r *Rectangle) drawHitbox(g *draw2dimg.GraphicContext) {
	var angle = r.getOffsetAngle()
	var position = r.getOffsetPosition()
	var width = r.width / 2
	var height = r.height / 2

	var topLeft = types.NewVector(-width, -height)
	var bottomRight = types.NewVector(width, height)

	g.Save()
	g.SetFillColor(r.collider.color)
	g.SetStrokeColor(r.collider.color)
	g.Translate(position.X, position.Y)
	g.Rotate(angle)
	g.BeginPath()
	draw2dkit.Rectangle(g, topLeft.X, topLeft.Y, bottomRight.X, bottomRight.Y)
	g.FillStroke()
	g.Restore()
}
