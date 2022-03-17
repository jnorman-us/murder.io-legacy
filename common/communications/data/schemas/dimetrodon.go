package schemas

import "github.com/josephnormandev/murder/common/communications/data"

var dimetrodonSchema = data.NewSchema(
	[]string{},
	[]string{"Health"},
	[]string{},
)
var DimetrodonSchema = &dimetrodonSchema

var dimetrodonA = data.MergeSchema(colliderSchema, dimetrodonSchema)
var dimetrodonStartSchema = data.MergeSchema(dimetrodonA, data.NewSchema(
	[]string{},
	[]string{},
	[]string{"Username"},
))
var DimetrodonStartSchema = &dimetrodonStartSchema
