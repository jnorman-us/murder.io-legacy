package world

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/types"
)

type Deletions struct {
	world   *World
	deleted map[types.ID]int
	flushed map[types.ID]int
}

func NewDeletions(w *World) *Deletions {
	return &Deletions{
		world:   w,
		deleted: map[types.ID]int{},
		flushed: map[types.ID]int{},
	}
}

func (d *Deletions) DeleteID(id types.ID) {
	d.deleted[id] = 0
}

func (d *Deletions) GetChannel() byte {
	return 0x01
}

func (d *Deletions) Flush() {
	d.flushed = d.deleted
	d.deleted = map[types.ID]int{}
}

func (d *Deletions) GetData(encoder *gob.Encoder) error {
	var flushed = d.flushed
	return encoder.Encode(flushed)
}

func (d *Deletions) HandleData(decoder *gob.Decoder) error {
	var deleted = &map[types.ID]int{}

	err := decoder.Decode(deleted)
	if err != nil {
		return err
	}

	for id := range *deleted {
		if _, ok := d.world.Poles[id]; ok {
			d.world.RemovePole(id)
		}
		if _, ok := d.world.Bullets[id]; ok {
			d.world.RemoveBullet(id)
		}
		if _, ok := d.world.Drifters[id]; ok {
			d.world.RemoveDrifter(id)
		}
	}
	return nil
}
