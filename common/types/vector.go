package types

import (
	"math"
	"math/rand"
)

type Vector struct {
	X, Y float64
}

func NewZeroVector() Vector {
	return Vector{
		X: 0,
		Y: 0,
	}
}

func NewRandomVector(minX, minY, maxX, maxY float64) Vector {
	return Vector{
		X: rand.Float64()*(maxX-minX) + minX,
		Y: rand.Float64()*(maxY-minX) + minY,
	}
}

func NewVector(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func (v *Vector) Copy() Vector {
	return Vector{
		X: v.X,
		Y: v.Y,
	}
}

func (v *Vector) Add(o Vector) {
	v.X += o.X
	v.Y += o.Y
}

func (v *Vector) Rotate(a float64) {
	v.X = math.Cos(a) * v.X
	v.Y = math.Sin(a) * v.Y
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
