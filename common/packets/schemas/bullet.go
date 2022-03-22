package schemas

import (
	"github.com/josephnormandev/murder/common/packets"
)

var bulletSchema = packets.NewSchema(
	[]string{},
	[]string{},
	[]string{},
)
var BulletSchema = packets.MergeSchema(ColliderSchema, bulletSchema)
