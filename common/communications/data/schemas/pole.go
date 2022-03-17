package schemas

import "github.com/josephnormandev/murder/common/communications/data"

var poleSchema = data.NewSchema(
	[]string{},
	[]string{},
	[]string{},
)
var PoleSchema = &poleSchema

var poleStartSchema = data.MergeSchema(colliderSchema, poleSchema)
var PoleStartSchema = &poleStartSchema
