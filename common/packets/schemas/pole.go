package schemas

var PoleSchema = MergeSchema(ColliderSchema, NewSchema(
	0x83,
	[]string{},
	[]string{},
	[]string{},
))
