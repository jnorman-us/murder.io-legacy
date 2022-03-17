package schemas

import "github.com/josephnormandev/murder/common/communications/data"

var drifterSchema = data.NewSchema(
	[]string{},
	[]string{"Health"},
	[]string{},
)
var DrifterSchema = &drifterSchema

var drifterA = data.MergeSchema(colliderSchema, drifterSchema)
var drifterStartSchema = data.MergeSchema(drifterA, data.NewSchema(
	[]string{},
	[]string{},
	[]string{"Username"},
))
var DrifterStartSchema = &drifterStartSchema
