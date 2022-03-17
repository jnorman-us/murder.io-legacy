package engine

import (
	"fmt"
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/communications/data/schemas"
	"github.com/josephnormandev/murder/common/types"
)

type Kinetic struct {
	types.ID

	StartPosition types.Vector
	StartAngle    float64

	EndPosition types.Vector
	EndAngle    float64

	started bool
}

func NewKinetic(id types.ID) *Kinetic {
	return &Kinetic{
		ID:      id,
		started: false,
	}
}

func GetDataID(datum data.Data) types.ID {
	datum.ApplySchema(schemas.MovementSchema)
	return types.ID(datum.GetInteger("ID"))
}

func (k *Kinetic) Set(position types.Vector, angle float64) {
	fmt.Println(k.StartPosition)
	if !k.started {
		k.StartPosition = position
		k.StartAngle = angle
		k.started = true
	} else {
		k.EndPosition = position
		k.EndAngle = angle
	}
	fmt.Println(k.StartPosition)
}

func (k *Kinetic) Restart() {
	fmt.Println("RESTARTING")
	k.StartPosition = k.EndPosition
	k.StartAngle = k.EndAngle
	k.started = false
}

func (k *Kinetic) Moved() bool {
	return !k.StartPosition.Equals(k.EndPosition) || k.StartAngle != k.EndAngle
}

func (k *Kinetic) GetData() data.Data {
	var datum = data.NewData(schemas.MovementSchema)
	datum.SetFloat("StartX", k.StartPosition.X)
	datum.SetFloat("StartY", k.StartPosition.Y)
	datum.SetFloat("StartAngle", k.StartAngle)
	datum.SetFloat("EndX", k.EndPosition.X)
	datum.SetFloat("EndY", k.EndPosition.Y)
	datum.SetFloat("EndAngle", k.EndAngle)
	datum.SetInteger("ID", int(k.ID))
	return datum
}

func (k *Kinetic) FromData(datum data.Data) {
	datum.ApplySchema(schemas.MovementSchema)

	k.StartPosition = types.NewVector(
		datum.GetFloat("StartX"),
		datum.GetFloat("StartY"),
	)
	k.StartAngle = datum.GetFloat("StartAngle")
	k.EndPosition = types.NewVector(
		datum.GetFloat("EndX"),
		datum.GetFloat("EndY"),
	)
	k.EndAngle = datum.GetFloat("EndAngle")
}
