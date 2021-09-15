package classes

import "github.com/josephnormandev/murder/types"

type Moveable interface {
	UpdatePosition()

	GetPosition() types.Vector
	SetPosition(types.Vector)
	GetAngle() float64
	SetAngle(float64)

	SetVelocity(types.Vector)
	GetVelocity() types.Vector
	SetAngularVelocity(float64)
	GetAngularVelocity() float64
}
