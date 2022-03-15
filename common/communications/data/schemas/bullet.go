package schemas

import "github.com/josephnormandev/murder/common/communications/data"

var bulletSchema = data.MergeSchema(colliderSchema, data.NewSchema(
	[]string{},
	[]string{},
	[]string{},
))
var BulletSchema = &bulletSchema
