package world

import (
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/packets"
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

func (w *World) AddByData(class types.Channel, data packets.Data) {
	var id = data.GetID()
	switch class {
	case dimetrodon.Class:
		var newDimetrodon = dimetrodon.NewDimetrodon()
		newDimetrodon.ID = id
		newDimetrodon.FromData(data)
		w.AddDimetrodon(newDimetrodon)
	case pole.Class:
		var newPole = pole.NewPole()
		newPole.ID = id
		newPole.FromData(data)
		w.AddPole(newPole)
	case bullet.Class:
		var newBullet = bullet.NewBullet(nil, 0)
		newBullet.ID = id
		newBullet.FromData(data)
		w.AddBullet(newBullet)
	}
}

func (w *World) UpdateByData(class types.Channel, data packets.Data) {
	var id = data.GetID()
	switch class {
	case dimetrodon.Class:
		if oldDimetrodon, ok := w.Dimetrodons[id]; ok {
			oldDimetrodon.FromData(data)
		}
	case pole.Class:
		if oldPole, ok := w.Poles[id]; ok {
			oldPole.FromData(data)
		}
	case bullet.Class:
		if oldBullet, ok := w.Bullets[id]; ok {
			oldBullet.FromData(data)
		}
	}
}

func (w *World) RemoveByData(class types.Channel, data packets.Data) {
	var id = data.GetID()
	switch class {
	case dimetrodon.Class:
		if oldDimetrodon, ok := w.Dimetrodons[id]; ok {
			oldDimetrodon.FromData(data)
			w.RemoveDimetrodon(id)
		}
	case pole.Class:
		if oldPole, ok := w.Poles[id]; ok {
			oldPole.FromData(data)
			w.RemovePole(id)
		}
	case bullet.Class:
		if oldBullet, ok := w.Bullets[id]; ok {
			oldBullet.FromData(data)
			w.RemoveBullet(id)
		}
	}
}
