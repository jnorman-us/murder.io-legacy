package schemas

var DimetrodonSchema = MergeSchema(ColliderSchema, NewSchema(
	0x82,
	[]string{},
	[]string{"Health", "MaxHealth"},
	[]string{"Username"},
))
