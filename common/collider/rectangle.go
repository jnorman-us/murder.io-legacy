package collider

import (
	"github.com/josephnormandev/murder/common/types"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

type Rectangle struct {
	localPosition types.Vector
	localAngle    float64
	width         float64
	height        float64
	collider      *Collider
	colliding     bool
}

func NewRectangle(p types.Vector, a, w, h float64) Rectangle {
	return Rectangle{
		localPosition: p,
		localAngle:    a,
		width:         w,
		height:        h,
	}
}

func (r *Rectangle) setCollider(c *Collider) {
	r.collider = c
}

func (r *Rectangle) checkCircleCollision(o *Circle) bool {
	return o.checkRectangleCollision(r)
}

func (r *Rectangle) checkRectangleCollision(o *Rectangle) bool {
	return false
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

func (r *Rectangle) draw(g *draw2dimg.GraphicContext) {
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
