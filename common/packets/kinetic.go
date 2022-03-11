package packets

import (
	"github.com/josephnormandev/murder/common/types"
	"time"
)

// Kinetic is a struct containing data about the Collider that is
// used for historical data as well as transporting over ws
type Kinetic struct {
	new        bool
	Offset     time.Duration
	StartX     float32
	StartY     float32
	StartAngle float32
	EndX       float32
	EndY       float32
	EndAngle   float32
}

func (k *Kinetic) SetData(position types.Vector, angle float64) {
	if k.new {
		k.StartX = float32(position.X)
		k.StartY = float32(position.Y)
		k.StartAngle = float32(angle)
		k.new = false
	} else {
		k.EndX = float32(position.X)
		k.EndY = float32(position.Y)
		k.EndAngle = float32(angle)
	}
}

func (k *Kinetic) Reset() {
	k.Offset = 0
	k.new = true
}

func (k *Kinetic) Moved() bool {
	return k.StartX != k.EndX || k.StartY != k.EndY || k.StartAngle != k.EndAngle
}

func (k *Kinetic) GetStartPosition() types.Vector {
	return types.NewVector(float64(k.StartX), float64(k.StartY))
}

func (k *Kinetic) GetEndPosition() types.Vector {
	return types.NewVector(float64(k.EndX), float64(k.EndY))
}

func (k *Kinetic) GetStartAngle() float64 {
	return float64(k.StartAngle)
}

func (k *Kinetic) GetEndAngle() float64 {
	return float64(k.EndAngle)
}
