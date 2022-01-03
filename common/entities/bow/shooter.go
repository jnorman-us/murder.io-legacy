package bow

import "github.com/josephnormandev/murder/common/types"

type Shooter interface {
	GetAngle() float64
	ScaleMass(float64)
	ResetMass()
	GetPosition() types.Vector
	GetVelocity() types.Vector
	GetFriction() float64
}
