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
