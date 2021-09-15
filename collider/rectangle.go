package collider

import (
	"github.com/josephnormandev/murder/types"
)

type Rectangle struct {
	localPosition types.Vector
	localAngle    float64
	width         float64
	height        float64

	color     types.Color
	collider  *Collider
	colliding bool
}

func NewRectangle(p types.Vector, a, w, h float64) Rectangle {
	return Rectangle{
		localPosition: p,
		localAngle:    a,
		width:         w,
		height:        h,
		color:         types.RandomColor(),
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

func (r *Rectangle) Draw(setCell func(types.Vector, rune)) {
	var angle = r.getOffsetAngle()
	var position = r.getOffsetPosition()

	for y := r.height / -2; y < r.height/2; y += 1 {
		for x := r.width / -2; x < r.width/2; x += 1 {
			var cell = types.NewVector(x+position.X, y+position.Y)
			cell.RotateAbout(angle, position)

			setCell(
				cell,
				'#',
			)
		}
	}
}
