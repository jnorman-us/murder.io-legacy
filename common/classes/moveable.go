package classes

import (
	types2 "github.com/josephnormandev/murder/common/types"
)

type Moveable interface {
	UpdatePosition()

	GetPosition() types2.Vector
	SetPosition(types2.Vector)
	GetAngle() float64
	SetAngle(float64)

	SetVelocity(types2.Vector)
	GetVelocity() types2.Vector
	SetAngularVelocity(float64)
	GetAngularVelocity() float64
}
