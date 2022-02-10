package logic

import (
	"github.com/josephnormandev/murder/common/types"
	"math"
)

// Driveable is a shared logic interface, allowing the same driving
// logic to be applied to different driveable Cars
type Driveable interface {
	GetTurningFactor() float64
	GetDrivingForce() float64
	GetDriftingFactor() float64
	GetDriftingReduction() float64
	GetFrontOfCar() types.Vector
	GetRearOfCar() types.Vector
	GetInput() types.Input
	GetAngle() float64
	GetVelocity() types.Vector
	GetPosition() types.Vector
	ApplyPositionalForce(types.Vector, types.Vector)
	ApplyPositionalForceAround(types.Vector, types.Vector, types.Vector)
}

func (m *Manager) AddDriveable(id types.ID, d *Driveable) {
	m.Driveables[id] = d
}

func (m *Manager) RemoveDriveable(id types.ID) {
	delete(m.Driveables, id)
}

func (m *Manager) Drive(d *Driveable) {
	var driveable = *d
	var input = driveable.GetInput()

	var frontOfCar = driveable.GetFrontOfCar()
	var rearOfCar = driveable.GetRearOfCar()

	var drivingForce = types.NewVector(driveable.GetDrivingForce(), 0)
	drivingForce.RotateAbout(driveable.GetAngle(), types.NewZeroVector())
	if input.Forward {
		drivingForce.Scale(1)
	} else if input.Backward {
		drivingForce.Scale(-1)
	} else if input.Left || input.Right {
		drivingForce.Scale(.5)
	} else {
		drivingForce.Scale(0)
	}
	if input.Left {
		drivingForce.RotateAbout(-1*driveable.GetTurningFactor(), types.NewZeroVector())
	} else if input.Right {
		drivingForce.RotateAbout(driveable.GetTurningFactor(), types.NewZeroVector())
	}

	var velocity = driveable.GetVelocity()
	var speed = velocity.Magnitude()

	var driftingForce = velocity
	driftingForce.Scale(math.Pow(speed, 1/2))
	driftingForce.Scale(driveable.GetDriftingFactor())

	if !input.Special {
		driftingForce.Scale(0)
	} else {
		drivingForce.Scale(driveable.GetDriftingReduction())
	}

	driveable.ApplyPositionalForceAround(drivingForce, frontOfCar, rearOfCar)
	driveable.ApplyPositionalForceAround(driftingForce, rearOfCar, frontOfCar)

	// TO DO, driving logic

	// parse the input

	// apply forces at nose, rear

	// consider drifting
}
