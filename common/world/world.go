package world

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

type World struct {
	spawner     *Spawner
	additions   *Additions
	deletions   *Deletions
	Drifters    map[types.ID]*drifter.Drifter // cars
	Dimetrodons map[types.ID]*dimetrodon.Dimetrodon
	Poles       map[types.ID]*pole.Pole // terrain elements
	Bullets     map[types.ID]*bullet.Bullet
}

func NewWorld(s *Spawner) *World {
	var game = &World{
		spawner:     s,
		Drifters:    map[types.ID]*drifter.Drifter{},
		Dimetrodons: map[types.ID]*dimetrodon.Dimetrodon{},
		Poles:       map[types.ID]*pole.Pole{},
		Bullets:     map[types.ID]*bullet.Bullet{},
	}
	return game
}

func (w *World) SetAdditions(additions *Additions) {
	w.additions = additions
}

func (w *World) SetDeletions(deletions *Deletions) {
	w.deletions = deletions
}

func (w *World) GetAdditions() *Additions {
	return w.additions
}

func (w *World) GetDeletions() *Deletions {
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
		} else { // world
		}
		break
	case dimetrodon.Class:
		var newDimetrodon = &dimetrodon.Dimetrodon{}

		err := decoder.Decode(newDimetrodon)
		if err != nil {
			return err
		}

		var _, ok = w.Dimetrodons[id]
		if !ok { // new, so add it
			w.AddDimetrodon(newDimetrodon)
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
		dimetrodon.Class,
		pole.Class,
		bullet.Class,
	}
}
