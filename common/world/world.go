package world

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

type World struct {
	spawner   *Spawner
	deletions *Deletions
	Drifters  map[types.ID]*drifter.Drifter // cars
	Poles     map[types.ID]*pole.Pole       // terrain elements
}

func NewWorld() *World {
	var game = &World{
		deletions: NewDeletions(),
		Drifters:  map[types.ID]*drifter.Drifter{},
		Poles:     map[types.ID]*pole.Pole{},
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
	}
	return nil
}

func (w *World) GetClasses() []byte {
	return []byte{
		drifter.Class,
		pole.Class,
	}
}
