package world

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/types"
)

type Deletions struct {
	deleted map[types.ID]int
	flushed map[types.ID]int
}

func NewDeletions() *Deletions {
	return &Deletions{
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
	return encoder.Encode(d.flushed)
}

func (d *Deletions) GetDeletions() []types.ID {
	d.Flush()
	var ids []types.ID
	for id := range d.flushed {
		ids = append(ids, id)
	}
	return ids
}

func (d *Deletions) HandleData(decoder *gob.Decoder) error {
	var deleted = &map[types.ID]int{}

	err := decoder.Decode(deleted)
	if err != nil {
		return err
	}

	for id := range *deleted {
		d.deleted[id] = 0
	}
	return nil
}
