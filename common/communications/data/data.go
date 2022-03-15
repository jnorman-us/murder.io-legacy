package data

import (
	"time"
)

type Data struct {
	RawFloats   []float32
	RawIntegers []int32
	RawStrings  []string

	offset time.Duration

	schema *Schema
}

func NewData(schema *Schema) Data {
	return Data{
		RawFloats:   make([]float32, len(schema.nameToFloat)),
		RawIntegers: make([]int32, len(schema.nameToInteger)),
		RawStrings:  make([]string, len(schema.nameToString)),

		offset: time.Duration(0),

		schema: schema,
	}
}

func (d *Data) ApplySchema(schema *Schema) {
	d.schema = schema
}

func (d *Data) SetOffset(offset time.Duration) {
	d.offset = offset
}

func (d *Data) GetOffset() time.Duration {
	return d.offset
}

func (d *Data) GetFloat(name string) float64 {
	if d.schema == nil {
		return 0.0
	}
	if index, ok := d.schema.nameToFloat[name]; ok {
		return float64(d.RawFloats[index])
	}
	return 0.0
}

func (d *Data) GetInteger(name string) int {
	if d.schema == nil {
		return 0
	}
	if index, ok := d.schema.nameToInteger[name]; ok {
		return int(d.RawIntegers[index])
	}
	return 0
}

func (d *Data) GetString(name string) string {
	if d.schema == nil {
		return ""
	}
	if index, ok := d.schema.nameToString[name]; ok {
		return d.RawStrings[index]
	}
	return ""
}

func (d *Data) SetFloat(name string, set float64) {
	if d.schema == nil {
		return
	}
	if index, ok := d.schema.nameToFloat[name]; ok {
		d.RawFloats[index] = float32(set)
	}
}

func (d *Data) SetInteger(name string, set int) {
	if d.schema == nil {
		return
	}
	if index, ok := d.schema.nameToInteger[name]; ok {
		d.RawIntegers[index] = int32(set)
	}
}

func (d *Data) SetString(name string, set string) {
	if d.schema == nil {
		return
	}
	if index, ok := d.schema.nameToString[name]; ok {
		d.RawStrings[index] = set
	}
}
