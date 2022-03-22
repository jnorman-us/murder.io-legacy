package schemas

import (
	"github.com/josephnormandev/murder/common/packets"
)

var ColliderSchema = packets.NewSchema(
	[]string{"X", "Y", "Angle"},
	[]string{},
	[]string{},
)
