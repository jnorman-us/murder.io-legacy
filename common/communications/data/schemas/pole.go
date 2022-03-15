package schemas

import "github.com/josephnormandev/murder/common/communications/data"

var poleSchema = data.MergeSchema(colliderSchema, data.NewSchema(
	[]string{},
	[]string{},
	[]string{},
))
var PoleSchema = &poleSchema
