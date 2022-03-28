package packets

import (
	"fmt"
	"github.com/josephnormandev/murder/common/packets/schemas"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/types/action"
	"github.com/josephnormandev/murder/common/types/timestamp"
)

type Data struct {
	timestamp *timestamp.Timestamp

	types.ID

	RawFloats   map[byte]float32
	RawIntegers map[byte]int32
	RawStrings  map[byte]string

	types.Change
	floatDiff  chan FloatDiff
	intDiff    chan IntDiff
	stringDiff chan StringDiff

	floatI   int
	intI     int
	stringI  int
	floatIn  []FloatDiff
	intIn    []IntDiff
	stringIn []StringDiff

	channel types.Channel
	schema  schemas.Schema
}

func NewData(id types.ID, c types.Channel, t *timestamp.Timestamp) *Data {
	return &Data{
		timestamp: t,

		ID: id,

		RawFloats:   map[byte]float32{},
		RawIntegers: map[byte]int32{},
		RawStrings:  map[byte]string{},

		floatDiff:  make(chan FloatDiff, 1000),
		intDiff:    make(chan IntDiff, 1000),
		stringDiff: make(chan StringDiff, 1000),

		channel: c,
		schema:  schemas.Schemas[c],
	}
}

func (d *Data) ApplySchema(schema schemas.Schema) {
	d.schema = schema
}

func (d *Data) GetFloat(name string) float64 {
	if !d.schema.Real() {
		return 0.0
	}
	if index, ok := d.schema.FloatNTI[name]; ok {
		return float64(d.RawFloats[index])
	}
	return 0.0
}

func (d *Data) GetInteger(name string) int {
	if !d.schema.Real() {
		return 0
	}
	if index, ok := d.schema.IntegerNTI[name]; ok {
		return int(d.RawIntegers[index])
	}
	return 0
}

func (d *Data) GetString(name string) string {
	if !d.schema.Real() {
		return ""
	}
	if index, ok := d.schema.StringNTI[name]; ok {
		return d.RawStrings[index]
	}
	return ""
}

func (d *Data) SetFloat(name string, set float64) {
	if !d.schema.Real() {
		return
	}
	if index, ok := d.schema.FloatNTI[name]; ok {
		var value = float32(set)
		var old, ok = d.RawFloats[index]
		if !ok || old != value {
			d.MarkDirty()
			d.RawFloats[index] = value
			d.floatDiff <- FloatDiff{
				Field:  index,
				Offset: d.timestamp.GetOffsetBytes(),
				Value:  value,
			}
		}
	}
}

func (d *Data) SetInteger(name string, set int) {
	if !d.schema.Real() {
		return
	}
	if index, ok := d.schema.IntegerNTI[name]; ok {
		var value = int32(set)
		var old, ok = d.RawIntegers[index]
		if !ok || old != value {
			d.MarkDirty()
			d.RawIntegers[index] = value
			d.intDiff <- IntDiff{
				Field:  index,
				Offset: d.timestamp.GetOffsetBytes(),
				Value:  value,
			}
		}
	}
}

func (d *Data) SetString(name string, set string) {
	if !d.schema.Real() {
		return
	}
	if index, ok := d.schema.StringNTI[name]; ok {
		var old, ok = d.RawStrings[index]
		if !ok || old != set {
			d.MarkDirty()
			d.RawStrings[index] = set
			d.stringDiff <- StringDiff{
				Field:  index,
				Offset: d.timestamp.GetOffsetBytes(),
				Value:  set,
			}
		}
	}
}

func (d *Data) GetChannel() types.Channel {
	return d.channel
}

func (d *Data) GenerateFullPacket() Packet {
	var floatDiff []FloatDiff
	var intDiff []IntDiff
	var stringDiff []StringDiff

	// generate full diffs
	for index, value := range d.RawFloats {
		floatDiff = append(floatDiff, FloatDiff{
			Field:  index,
			Offset: 0,
			Value:  value,
		})
	}
	for index, value := range d.RawIntegers {
		intDiff = append(intDiff, IntDiff{
			Field:  index,
			Offset: 0,
			Value:  value,
		})
	}
	for index, value := range d.RawStrings {
		stringDiff = append(stringDiff, StringDiff{
			Field:  index,
			Offset: 0,
			Value:  value,
		})
	}

	return Packet{
		ID:         d.ID,
		Channel:    d.channel,
		Action:     action.Actions.Add,
		Offset:     0,
		FloatDiff:  floatDiff,
		IntDiff:    intDiff,
		StringDiff: stringDiff,
	}
}

func (d *Data) GeneratePacket(a action.Action, offset byte) Packet {
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
				ID:         d.ID,
				Channel:    d.channel,
				Action:     a,
				Offset:     offset,
				FloatDiff:  floatDiff,
				IntDiff:    intDiff,
				StringDiff: stringDiff,
			}
		}
	}
}

func (d *Data) SetDiffs(floatDiff []FloatDiff, intDiff []IntDiff, stringDiff []StringDiff) {
	d.floatI = 0
	d.intI = 0
	d.stringI = 0
	d.floatIn = floatDiff
	d.intIn = intDiff
	d.stringIn = stringDiff
}

func (d *Data) Trickle(elapsed byte) bool {
	var trickled = false
	for ; d.floatI < len(d.floatIn); d.floatI++ {
		var floatDiff = d.floatIn[d.floatI]
		if floatDiff.Offset <= elapsed {
			d.applyFloatDiff(floatDiff)
			trickled = true
		} else {
			break
		}
	}
	for ; d.intI < len(d.intIn); d.intI++ {
		var intDiff = d.intIn[d.intI]
		if intDiff.Offset <= elapsed {
			d.applyIntegerDiff(intDiff)
			trickled = true
		} else {
			break
		}
	}
	for ; d.stringI < len(d.stringIn); d.stringI++ {
		var stringDiff = d.stringIn[d.stringI]
		if stringDiff.Offset <= elapsed {
			d.applyStringDiff(stringDiff)
			trickled = true
		} else {
			break
		}
	}
	return trickled
}

func (d *Data) applyFloatDiff(diff FloatDiff) {
	var field = diff.Field
	var value = diff.Value
	d.RawFloats[field] = value
}

func (d *Data) applyIntegerDiff(diff IntDiff) {
	var field = diff.Field
	var value = diff.Value
	d.RawIntegers[field] = value
}

func (d *Data) applyStringDiff(diff StringDiff) {
	var field = diff.Field
	var value = diff.Value
	d.RawStrings[field] = value
}

func (d *Data) Print() {
	fmt.Println(d.ID, d.RawFloats, d.RawIntegers, d.RawStrings)
}
