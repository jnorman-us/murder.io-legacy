package drifter

import "github.com/josephnormandev/murder/common/types"

func (d *Drifter) GetTurningFactor() float64 {
	return .0002
}

func (d *Drifter) GetDrivingForce() float64 {
	return 220
}

func (d *Drifter) GetDriftingFactor() float64 {
	return .02
}

func (d *Drifter) GetDriftingReduction() float64 {
	return .7
}

func (d *Drifter) GetFrontOfCar() types.Vector {
	var position = d.GetPosition()
	var lengthOfCar = types.NewVector(22.5, 0)
	lengthOfCar.RotateAbout(d.GetAngle(), types.NewZeroVector())
	position.Add(lengthOfCar)
	return position
}

func (d *Drifter) GetRearOfCar() types.Vector {
	var position = d.GetPosition()
	var lengthOfCar = types.NewVector(-22.5, 0)
	lengthOfCar.RotateAbout(d.GetAngle(), types.NewZeroVector())
	position.Add(lengthOfCar)
	return position
}
