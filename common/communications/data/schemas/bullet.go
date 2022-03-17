package schemas

import "github.com/josephnormandev/murder/common/communications/data"

var bulletSchema = data.NewSchema(
	[]string{},
	[]string{},
	[]string{},
)
var BulletSchema = &bulletSchema

var bulletStartSchema = data.MergeSchema(colliderSchema, bulletSchema)
var BulletStartSchema = &bulletStartSchema
