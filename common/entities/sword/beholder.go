package sword

import "github.com/josephnormandev/murder/common/types"

type Beholder interface {
	GetID() int
	GetIdentifier() string
	GetPosition() types.Vector
	GetVelocity() types.Vector
	GetFriction() float64
	GetAngle() float64
}
