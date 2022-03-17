package world

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

type World struct {
	spawner *Spawner

	Dimetrodons map[types.ID]*dimetrodon.Dimetrodon
	Poles       map[types.ID]*pole.Pole // terrain elements
	Bullets     map[types.ID]*bullet.Bullet
}

func NewWorld(s *Spawner) *World {
	var game = &World{
		spawner:     s,
		Dimetrodons: map[types.ID]*dimetrodon.Dimetrodon{},
		Poles:       map[types.ID]*pole.Pole{},
		Bullets:     map[types.ID]*bullet.Bullet{},
	}
	return game
}

func (w *World) HandleAddition(id types.ID, channel types.Channel, datum data.Data) {
	switch channel {
	case dimetrodon.Class:
		d := dimetrodon.NewDimetrodon()
		d.ID = id
		d.FromStartData(datum)
		w.AddDimetrodon(d)
		break
	case bullet.Class:
		b := bullet.NewBullet(nil, 0)
		b.ID = id
		b.FromStartData(datum)
		w.AddBullet(b)
		break
	case pole.Class:
		p := pole.NewPole()
		p.ID = id
		p.FromStartData(datum)
		w.AddPole(p)
		break
	}
}

func (w *World) HandleUpdate(id types.ID, channel types.Channel, datum data.Data) {
	switch channel {
	case dimetrodon.Class:
		if d, ok := w.Dimetrodons[id]; ok {
			d.FromData(datum)
		}
		break
	case bullet.Class:
		if d, ok := w.Bullets[id]; ok {
			d.FromData(datum)
		}
		break
	case pole.Class:
		if d, ok := w.Poles[id]; ok {
			d.FromData(datum)
		}
		break
	}
}

func (w *World) HandleDeletion(id types.ID, channel types.Channel, datum data.Data) {
	switch channel {
	case dimetrodon.Class:
		w.RemoveDimetrodon(id)
		break
	case bullet.Class:
		w.RemoveBullet(id)
		break
	case pole.Class:
		w.RemovePole(id)
		break
	}
}
