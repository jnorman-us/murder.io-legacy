package interpolater

import "github.com/josephnormandev/murder/common/types"

type Interpolatable interface {
	GetID() int

	SetPosition(types.Vector)
	SetVelocity(types.Vector)
	SetAngle(float64)

	GetPosition() types.Vector
	GetVelocity() types.Vector
	GetAngle() float64
}
