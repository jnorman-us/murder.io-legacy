package world

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"time"
)

type Additions struct {
	world   *World
	time    *types.Time
	added   map[types.ID]packets.Addition
	flushed []packets.Addition
}

func NewAdditions(w *World, t *types.Time) *Additions {
	return &Additions{
		world:   w,
		time:    t,
		added:   map[types.ID]packets.Addition{},
		flushed: []packets.Addition{},
	}
}

func (a *Additions) Add(id types.ID, class byte) {
	a.added[id] = packets.Addition{
		ID:     id,
		Class:  class,
		Offset: a.time.GetOffset(),
	}
}

func (a *Additions) GetChannel() byte {
	return 0x05
}

func (a *Additions) Flush() {
	a.flushed = []packets.Addition{}
	for _, addition := range a.added {
		a.flushed = append(a.flushed, addition)
	}
	a.added = map[types.ID]packets.Addition{}
}

func (a *Additions) GetData(encoder *gob.Encoder) error {
	var flushed = a.flushed
	return encoder.Encode(flushed)
}

func (a *Additions) HandleData(decoder *gob.Decoder) error {
	var added = &[]packets.Addition{}

	err := decoder.Decode(added)
	if err != nil {
		return err
	}

	for _, addition := range *added {
		a.added[addition.ID] = addition
	}

	return nil
}

func (a *Additions) AddTick(elapsed, total time.Duration) {
	// do something FUCK
}
