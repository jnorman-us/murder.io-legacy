package schemas

import (
	"github.com/josephnormandev/murder/common/packets"
)

var dimetrodonSchema = packets.NewSchema(
	[]string{},
	[]string{"Health", "MaxHealth"},
	[]string{"Username"},
)
var DimetrodonSchema = packets.MergeSchema(ColliderSchema, dimetrodonSchema)
