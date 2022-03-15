package data

type Schema struct {
	nameToFloat   map[string]int
	nameToInteger map[string]int
	nameToString  map[string]int
}

func NewSchema(floatNames []string, integerNames []string, stringNames []string) Schema {
	var schema = Schema{
		nameToFloat:   map[string]int{},
		nameToInteger: map[string]int{},
		nameToString:  map[string]int{},
	}

	for index, name := range floatNames {
		schema.nameToFloat[name] = index
	}
	for index, name := range integerNames {
		schema.nameToInteger[name] = index
	}
	for index, name := range stringNames {
		schema.nameToString[name] = index
	}

	return schema
}

func MergeSchema(a Schema, b Schema) Schema {
	var c = Schema{
		nameToFloat:   map[string]int{},
		nameToInteger: map[string]int{},
		nameToString:  map[string]int{},
	}

	for name, index := range a.nameToFloat {
		c.nameToFloat[name] = index
	}
	var floatLen = len(a.nameToFloat)
	for name, index := range b.nameToFloat {
		c.nameToFloat[name] = floatLen + index
	}

	for name, index := range a.nameToInteger {
		c.nameToInteger[name] = index
	}
	var integerLen = len(a.nameToInteger)
	for name, index := range b.nameToInteger {
		c.nameToInteger[name] = integerLen + index
	}

	for name, index := range a.nameToString {
		c.nameToString[name] = index
	}
	var stringLen = len(a.nameToString)
	for name, index := range b.nameToString {
		c.nameToString[name] = stringLen + index
	}

	return c
}
