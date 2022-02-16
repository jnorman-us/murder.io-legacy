package types

import (
	"github.com/Tarliton/collision2d"
	"math"
	"math/rand"
)

type Vector struct {
	collision2d.Vector
}

func NewZeroVector() Vector {
	return Vector{
		collision2d.NewVector(0, 0),
	}
}

func NewRandomVector(minX, minY, maxX, maxY float64) Vector {
	return Vector{
		collision2d.NewVector(
			rand.Float64()*(maxX-minX)+minX,
			rand.Float64()*(maxY-minX)+minY,
		),
	}
}

func NewVector(x, y float64) Vector {
	return Vector{
		collision2d.NewVector(x, y),
	}
}

func (v *Vector) Copy() Vector {
	return Vector{
		collision2d.Vector{
			X: v.X,
			Y: v.Y,
		},
	}
}

func (v *Vector) Add(o Vector) {
	v.X += o.X
	v.Y += o.Y
}

func (v *Vector) Interpolate(b Vector, alpha float64) {
	v.X += (b.X - v.X) * alpha
	v.Y += (b.Y - v.Y) * alpha
}

func (v *Vector) MultiplyBy(o Vector) {
	v.X *= o.X
	v.Y *= o.Y
}

func (v *Vector) Rotate(a float64) {
	v.X = math.Cos(a)*v.X - math.Sin(a)*v.Y
	v.Y = math.Sin(a)*v.X + math.Cos(a)*v.Y
}

func (v *Vector) RotateAbout(a float64, o Vector) {
	var cos = math.Cos(a)
	var sin = math.Sin(a)
	var vector = v.Copy()

	v.X = o.X + ((vector.X-o.X)*cos - (vector.Y-o.Y)*sin)
	v.Y = o.Y + ((vector.X-o.X)*sin + (vector.Y-o.Y)*cos)
}

func (v *Vector) Scale(scalar float64) {
	v.X = v.X * scalar
	v.Y = v.Y * scalar
}

func (v *Vector) Offset(o Vector) Vector {
	return Vector{
		collision2d.Vector{
			X: o.X - v.X,
			Y: o.Y - v.Y,
		},
	}
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Distance(o Vector) float64 {
	var x = v.X - o.X
	var y = v.Y - o.Y
	return math.Sqrt((x * x) + (y * y))
}

func (v *Vector) Angle() float64 {
	if v.X == 0 {
		return math.Atan(10)
	}
	var angle = math.Atan(v.Y / v.X)

	var output = angle
	if v.X <= 0 {
		output = output - math.Pi
	}
	return output
}

func (v *Vector) Equals(o Vector) bool {
	return v.X == o.X && v.Y == o.Y
}
