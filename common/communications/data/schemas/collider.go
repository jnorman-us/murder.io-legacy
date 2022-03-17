package schemas

import (
	"github.com/josephnormandev/murder/common/communications/data"
)

var colliderSchema = data.NewSchema(
	[]string{"X", "Y", "Angle"},
	[]string{},
	[]string{},
)
var ColliderSchema = &colliderSchema

var movementSchema = data.NewSchema(
	[]string{"StartX", "StartY", "StartAngle", "EndX", "EndY", "EndAngle"},
	[]string{"ID"},
	[]string{},
)
var MovementSchema = &movementSchema
