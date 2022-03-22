package packets

import (
	"fmt"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/types/action"
	"github.com/josephnormandev/murder/common/types/timestamp"
)

type Data struct {
	timestamp *timestamp.Timestamp

	id types.ID

	RawFloats   map[byte]float32
	RawIntegers map[byte]int32
	RawStrings  map[byte]string

	types.Change
	floatDiff  chan FloatDiff
	intDiff    chan IntDiff
	stringDiff chan StringDiff

	schema Schema
}

func newData(id types.ID, s Schema, t *timestamp.Timestamp) *Data {
	return &Data{
		timestamp: t,

		id: id,

		RawFloats:   map[byte]float32{},
		RawIntegers: map[byte]int32{},
		RawStrings:  map[byte]string{},

		floatDiff:  make(chan FloatDiff, 1000),
		intDiff:    make(chan IntDiff, 1000),
		stringDiff: make(chan StringDiff, 1000),

		schema: s,
	}
}

func (d *Data) ApplySchema(schema Schema) {
	d.schema = schema
}

func (d *Data) GetFloat(name string) float64 {
	if !d.schema.real {
		return 0.0
	}
	if index, ok := d.schema.floatNTI[name]; ok {
		return float64(d.RawFloats[index])
	}
	return 0.0
}

func (d *Data) GetInteger(name string) int {
	if !d.schema.real {
		return 0
	}
	if index, ok := d.schema.integerNTI[name]; ok {
		return int(d.RawIntegers[index])
	}
	return 0
}

func (d *Data) GetString(name string) string {
	if !d.schema.real {
		return ""
	}
	if index, ok := d.schema.stringNTI[name]; ok {
		return d.RawStrings[index]
	}
	return ""
}

func (d *Data) SetFloat(name string, set float64) {
	if !d.schema.real {
		return
	}
	if index, ok := d.schema.floatNTI[name]; ok {
		var value = float32(set)
		var old = d.RawFloats[index]
		if old != value {
			d.MarkDirty()
			d.RawFloats[index] = value
			fmt.Println(len(d.floatDiff), len(d.intDiff), len(d.stringDiff))
			d.floatDiff <- FloatDiff{
				field:  index,
				offset: d.timestamp.GetOffsetBytes(),
				value:  value,
			}
		}
	}
}

func (d *Data) SetInteger(name string, set int) {
	if !d.schema.real {
		return
	}
	if index, ok := d.schema.integerNTI[name]; ok {
		var value = int32(set)
		var old = d.RawIntegers[index]
		if old != value {
			d.MarkDirty()
			d.RawIntegers[index] = value
			d.intDiff <- IntDiff{
				field:  index,
				offset: d.timestamp.GetOffsetBytes(),
				value:  value,
			}
		}
	}
}

func (d *Data) SetString(name string, set string) {
	if !d.schema.real {
		return
	}
	if index, ok := d.schema.stringNTI[name]; ok {
		var old = d.RawStrings[index]
		if old != set {
			d.MarkDirty()
			d.RawStrings[index] = set
			d.stringDiff <- StringDiff{
				field:  index,
				offset: d.timestamp.GetOffsetBytes(),
				value:  set,
			}
		}
	}
}

func (d *Data) GenerateFullPacket(c types.Channel) Packet {
	var floatDiff []FloatDiff
	var intDiff []IntDiff
	var stringDiff []StringDiff

	// generate full diffs
	for index, value := range d.RawFloats {
		floatDiff = append(floatDiff, FloatDiff{
			field:  index,
			offset: 0,
			value:  value,
		})
	}
	for index, value := range d.RawIntegers {
		intDiff = append(intDiff, IntDiff{
			field:  index,
			offset: 0,
			value:  value,
		})
	}
	for index, value := range d.RawStrings {
		stringDiff = append(stringDiff, StringDiff{
			field:  index,
			offset: 0,
			value:  value,
		})
	}

	return Packet{
		ID:         d.id,
		Channel:    c,
		Action:     action.Actions.Add,
		Offset:     0,
		FloatDiff:  floatDiff,
		IntDiff:    intDiff,
		StringDiff: stringDiff,
	}
}

func (d *Data) GeneratePacket(c types.Channel, a action.Action, offset byte) Packet {
	var floatDiff []FloatDiff
	var intDiff []IntDiff
	var stringDiff []StringDiff

	for {
		select {
		case diff := <-d.floatDiff:
			floatDiff = append(floatDiff, diff)
		case diff := <-d.intDiff:
			intDiff = append(intDiff, diff)
		case diff := <-d.stringDiff:
			stringDiff = append(stringDiff, diff)
		default:
			return Packet{
				ID:         d.id,
				Channel:    c,
				Action:     a,
				Offset:     offset,
				FloatDiff:  floatDiff,
				IntDiff:    intDiff,
				StringDiff: stringDiff,
			}
		}
	}
}
