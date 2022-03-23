package schemas

var BulletSchema = MergeSchema(ColliderSchema, NewSchema(
	0x81,
	[]string{},
	[]string{},
	[]string{},
))
