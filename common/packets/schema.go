package packets

type Schema struct {
	real       bool
	floatNTI   map[string]byte
	integerNTI map[string]byte
	stringNTI  map[string]byte
}

func NewSchema(floatNames []string, integerNames []string, stringNames []string) Schema {
	var schema = Schema{
		real:       true,
		floatNTI:   map[string]byte{},
		integerNTI: map[string]byte{},
		stringNTI:  map[string]byte{},
	}

	for index, name := range floatNames {
		schema.floatNTI[name] = byte(index)
	}
	for index, name := range integerNames {
		schema.integerNTI[name] = byte(index)
	}
	for index, name := range stringNames {
		schema.stringNTI[name] = byte(index)
	}

	return schema
}

func MergeSchema(a Schema, b Schema) Schema {
	var c = Schema{
		real:       true,
		floatNTI:   map[string]byte{},
		integerNTI: map[string]byte{},
		stringNTI:  map[string]byte{},
	}

	for name, index := range a.floatNTI {
		c.floatNTI[name] = index
	}
	var floatLen = len(a.floatNTI)
	for name, index := range b.floatNTI {
		c.floatNTI[name] = byte(floatLen) + index
	}

	for name, index := range a.integerNTI {
		c.integerNTI[name] = index
	}
	var integerLen = len(a.integerNTI)
	for name, index := range b.integerNTI {
		c.integerNTI[name] = byte(integerLen) + index
	}

	for name, index := range a.stringNTI {
		c.stringNTI[name] = index
	}
	var stringLen = len(a.stringNTI)
	for name, index := range b.stringNTI {
		c.stringNTI[name] = byte(stringLen) + index
	}

	return c
}
