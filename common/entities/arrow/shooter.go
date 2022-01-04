package arrow

import "github.com/josephnormandev/murder/common/types"

type Shooter interface {
	GetID() int
	GetAngle() float64
	GetPosition() types.Vector
	GetVelocity() types.Vector
}
