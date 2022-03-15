package schemas

import "github.com/josephnormandev/murder/common/communications/data"

var dimetrodonSchema = data.MergeSchema(colliderSchema, data.NewSchema(
	[]string{},
	[]string{"Health", "MaxHealth"},
	[]string{"Username"},
))
var DimetrodonSchema = &dimetrodonSchema
