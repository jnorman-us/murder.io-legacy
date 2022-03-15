package schemas

import "github.com/josephnormandev/murder/common/communications/data"

var drifterSchema = data.MergeSchema(colliderSchema, data.NewSchema(
	[]string{},
	[]string{"Health", "MaxHealth"},
	[]string{"Username"},
))
var DrifterSchema = &drifterSchema
