package collider

import (
	"github.com/Tarliton/collision2d"
	"github.com/josephnormandev/murder/common/types"
)

type Rectangle struct {
	localPosition types.Vector
	localAngle    float64
	width         float64
	height        float64
	collider      *Collider

	calculatedPolygon collision2d.Polygon
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

func (r *Rectangle) calculate() {
	var width, height = r.width, r.height
	var angle = r.getOffsetAngle()
	var position = r.localPosition.Copy()
	position.Add(types.NewVector(width/-2, height/-2))
	position.Add(r.collider.GetPosition())
	position.RotateAbout(angle, r.collider.GetPosition())
	var box = collision2d.NewBox(collision2d.Vector(position), width, height)
	r.calculatedPolygon = box.ToPolygon()
}

func (r *Rectangle) getPolygon() collision2d.Polygon {
	return r.calculatedPolygon
}

func (r *Rectangle) drawHitbox() {
}
