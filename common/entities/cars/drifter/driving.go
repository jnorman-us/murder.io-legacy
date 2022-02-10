package drifter

import "github.com/josephnormandev/murder/common/types"

func (d *Drifter) GetTurningFactor() float64 {
	return .08
}

func (d *Drifter) GetDrivingForce() float64 {
	return 300
}

func (d *Drifter) GetDriftingFactor() float64 {
	return .1
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
