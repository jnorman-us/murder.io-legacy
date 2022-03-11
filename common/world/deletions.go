package world

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"time"
)

type Deletions struct {
	world   *World
	time    *types.Time
	deleted map[types.ID]packets.Deletion
	flushed []packets.Deletion
}

func NewDeletions(w *World, t *types.Time) *Deletions {
	return &Deletions{
		world:   w,
		time:    t,
		deleted: map[types.ID]packets.Deletion{},
		flushed: []packets.Deletion{},
	}
}

func (d *Deletions) Delete(id types.ID) {
	d.deleted[id] = packets.Deletion{
		ID:     id,
		Offset: d.time.GetOffset(),
	}
}

func (d *Deletions) GetChannel() byte {
	return 0x01
}

func (d *Deletions) Flush() {
	d.flushed = []packets.Deletion{}
	for _, deletion := range d.deleted {
		d.flushed = append(d.flushed, deletion)
	}
	d.deleted = map[types.ID]packets.Deletion{}
}

func (d *Deletions) GetData(encoder *gob.Encoder) error {
	var flushed = d.flushed
	return encoder.Encode(flushed)
}

// HandleFutureData takes the packet data and stores it until the DeleteTick
// clears it. If it is never cleared, then it will clear it on the next
// HandleData call
func (d *Deletions) HandleFutureData(decoder *gob.Decoder, ttl time.Duration) error {
	for id := range d.deleted {
		d.deleteID(id)
	}

	var deleted = &[]packets.Deletion{}

	err := decoder.Decode(deleted)
	if err != nil {
		return err
	}

	for _, dlt := range *deleted {
		d.deleted[dlt.ID] = dlt
	}

	return nil
}

// DeleteTick deletes all items in the map if it is their time
func (d *Deletions) DeleteTick(elapsed, total time.Duration) {
	for id, dl := range d.deleted {
		var offset = dl.Offset
		if offset <= elapsed {
			d.deleteID(id)
		}
	}
}

func (d *Deletions) deleteID(id types.ID) {
	delete(d.deleted, id)
	if _, ok := d.world.Bullets[id]; ok {
		d.world.RemoveBullet(id)
	}
	if _, ok := d.world.Drifters[id]; ok {
		d.world.RemoveDrifter(id)
	}
	if _, ok := d.world.Dimetrodons[id]; ok {
		d.world.RemoveDimetrodon(id)
	}
	if _, ok := d.world.Poles[id]; ok {
		d.world.RemovePole(id)
	}
}
