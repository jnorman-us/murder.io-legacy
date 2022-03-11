package world

import (
	"encoding/gob"
	"fmt"
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
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
	for _, addition := range a.added {
		a.flushed = append(a.flushed, addition)
	}
	a.added = map[types.ID]packets.Addition{}
}

func (a *Additions) GetData(encoder *gob.Encoder) error {
	var flushed = a.flushed
	a.flushed = []packets.Addition{}
	return encoder.Encode(flushed)
}

func (a *Additions) GetCatchupData(encoder *gob.Encoder) error {
	var world = a.world

	var added []packets.Addition
	for id := range world.Drifters {
		added = append(added, packets.Addition{
			ID:    id,
			Class: drifter.Class,
		})
	}
	for id := range world.Dimetrodons {
		added = append(added, packets.Addition{
			ID:    id,
			Class: dimetrodon.Class,
		})
	}
	for id := range world.Poles {
		added = append(added, packets.Addition{
			ID:    id,
			Class: pole.Class,
		})
	}
	for id := range world.Bullets {
		added = append(added, packets.Addition{
			ID:    id,
			Class: bullet.Class,
		})
	}

	return encoder.Encode(added)
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
	for id, add := range a.added {
		var class = add.Class
		var offset = add.Offset
		if offset <= elapsed {
			a.addID(id, class)
		}
	}
}

func (a *Additions) addID(id types.ID, class byte) {
	var world = a.world
	delete(a.added, id)

	switch class {
	case drifter.Class:
		var newDrifter = drifter.NewDrifter()
		newDrifter.ID = id
		if state, ok := world.DrifterTemp[id]; ok {
			newDrifter.State = state
			delete(world.DrifterTemp, id)
		}
		world.AddDrifter(newDrifter)
		fmt.Println("NEw Drifter", newDrifter)
		break
	case dimetrodon.Class:
		var newDimetrodon = dimetrodon.NewDimetrodon()
		newDimetrodon.ID = id
		if state, ok := world.DimetrodonTemp[id]; ok {
			newDimetrodon.State = state
			delete(world.DimetrodonTemp, id)
		}
		fmt.Println("New Dimetrodon", newDimetrodon)
		world.AddDimetrodon(newDimetrodon)
		break
	case pole.Class:
		var newPole = pole.NewPole()
		newPole.ID = id
		if state, ok := world.PoleTemp[id]; ok {
			newPole.State = state
			delete(world.PoleTemp, id)
		}
		fmt.Println(newPole)
		world.AddPole(newPole)
		break
	case bullet.Class:
		var newBullet = bullet.NewBullet(nil, 0)
		newBullet.ID = id
		if state, ok := world.BulletTemp[id]; ok {
			newBullet.State = state
			delete(world.BulletTemp, id)
		}
		world.AddBullet(newBullet)
		fmt.Println(newBullet)
		break
	}
}
