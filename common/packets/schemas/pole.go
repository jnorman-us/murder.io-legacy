package schemas

import (
	"github.com/josephnormandev/murder/common/packets"
)

var poleSchema = packets.NewSchema(
	[]string{},
	[]string{},
	[]string{},
)
var PoleSchema = packets.MergeSchema(ColliderSchema, poleSchema)
