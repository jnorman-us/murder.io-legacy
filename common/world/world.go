package world

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

type World struct {
	spawner   *Spawner
	deletions *Deletions
	Drifters  map[types.ID]*drifter.Drifter // cars
	Poles     map[types.ID]*pole.Pole       // terrain elements
	Bullets   map[types.ID]*bullet.Bullet
}

func NewWorld() *World {
	var game = &World{
		deletions: NewDeletions(),
		Drifters:  map[types.ID]*drifter.Drifter{},
		Poles:     map[types.ID]*pole.Pole{},
		Bullets:   map[types.ID]*bullet.Bullet{},
	}
	return game
}

func (w *World) SetSpawner(s *Spawner) {
	w.spawner = s
}

func (w *World) Deletions() *Deletions {
	return w.deletions
}

func (w *World) HandleSpawn(id types.ID, class byte, decoder *gob.Decoder) error {
	switch class {
	case drifter.Class:
		var newDrifter = &drifter.Drifter{}

		err := decoder.Decode(newDrifter)
		if err != nil {
			return err
		}

		var _, ok = w.Drifters[id]
		if !ok { // new, so add it
			w.AddDrifter(newDrifter)
		} else { // update
		}
		break
	case pole.Class:
		var newPole = &pole.Pole{}

		err := decoder.Decode(newPole)
		if err != nil {
			return err
		}

		var _, ok = w.Poles[id]
		if !ok {
			w.AddPole(newPole)
		} else { // update
		}
		break
	case bullet.Class:
		var newBullet = &bullet.Bullet{}

		err := decoder.Decode(newBullet)
		if err != nil {
			return err
		}
		var _, ok = w.Bullets[id]
		if !ok {
			w.AddBullet(newBullet)
		} else { // update
		}
		break
	}
	return nil
}

func (w *World) GetClasses() []byte {
	return []byte{
		drifter.Class,
		pole.Class,
		bullet.Class,
	}
}
