package schemas

import "github.com/josephnormandev/murder/common/types"

var Schemas = map[types.Channel]Schema{
	ColliderSchema.channel:   ColliderSchema,
	BulletSchema.channel:     BulletSchema,
	DimetrodonSchema.channel: DimetrodonSchema,
	PoleSchema.channel:       PoleSchema,
}

type Schema struct {
	real       bool
	channel    types.Channel
	FloatNTI   map[string]byte
	IntegerNTI map[string]byte
	StringNTI  map[string]byte
}

func NewSchema(c types.Channel, floatNames []string, integerNames []string, stringNames []string) Schema {
	var schema = Schema{
		real:       true,
		channel:    c,
		FloatNTI:   map[string]byte{},
		IntegerNTI: map[string]byte{},
		StringNTI:  map[string]byte{},
	}

	for index, name := range floatNames {
		schema.FloatNTI[name] = byte(index)
	}
	for index, name := range integerNames {
		schema.IntegerNTI[name] = byte(index)
	}
	for index, name := range stringNames {
		schema.StringNTI[name] = byte(index)
	}

	return schema
}

func (s *Schema) Channel() types.Channel {
	return s.channel
}

func (s *Schema) Real() bool {
	return s.real
}

func MergeSchema(a Schema, b Schema) Schema {
	var c = Schema{
		real:       true,
		channel:    b.channel,
		FloatNTI:   map[string]byte{},
		IntegerNTI: map[string]byte{},
		StringNTI:  map[string]byte{},
	}

	for name, index := range a.FloatNTI {
		c.FloatNTI[name] = index
	}
	var floatLen = len(a.FloatNTI)
	for name, index := range b.FloatNTI {
		c.FloatNTI[name] = byte(floatLen) + index
	}

	for name, index := range a.IntegerNTI {
		c.IntegerNTI[name] = index
	}
	var integerLen = len(a.IntegerNTI)
	for name, index := range b.IntegerNTI {
		c.IntegerNTI[name] = byte(integerLen) + index
	}

	for name, index := range a.StringNTI {
		c.StringNTI[name] = index
	}
	var stringLen = len(a.StringNTI)
	for name, index := range b.StringNTI {
		c.StringNTI[name] = byte(stringLen) + index
	}

	return c
}
