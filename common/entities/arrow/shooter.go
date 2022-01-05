package arrow

import "github.com/josephnormandev/murder/common/types"

type Shooter interface {
	GetID() int
	GetUsername() string
	GetAngle() float64
	GetPosition() types.Vector
	GetVelocity() types.Vector
}
