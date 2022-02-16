package dimetrodon

import "github.com/josephnormandev/murder/common/types"

func (d *Dimetrodon) GetTurningFactor() float64 {
	return .0002
}

func (d *Dimetrodon) GetDrivingForce() float64 {
	return 220
}

func (d *Dimetrodon) GetDriftingFactor() float64 {
	return .02
}

func (d *Dimetrodon) GetDriftingReduction() float64 {
	return .7
}

func (d *Dimetrodon) GetFrontOfCar() types.Vector {
	var position = d.GetPosition()
	var lengthOfCar = types.NewVector(22.5, 0)
	lengthOfCar.RotateAbout(d.GetAngle(), types.NewZeroVector())
	position.Add(lengthOfCar)
	return position
}

func (d *Dimetrodon) GetRearOfCar() types.Vector {
	var position = d.GetPosition()
	var lengthOfCar = types.NewVector(-22.5, 0)
	lengthOfCar.RotateAbout(d.GetAngle(), types.NewZeroVector())
	position.Add(lengthOfCar)
	return position
}
