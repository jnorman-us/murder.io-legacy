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
	environment types.Environment
	additions   *Additions
	deletions   *Deletions

	Drifters    map[types.ID]*drifter.Drifter // cars
	Dimetrodons map[types.ID]*dimetrodon.Dimetrodon
	Poles       map[types.ID]*pole.Pole // terrain elements
	Bullets     map[types.ID]*bullet.Bullet

	// state maps for each type. These are only used upon addition
	DrifterTemp    map[types.ID]drifter.State
	DimetrodonTemp map[types.ID]dimetrodon.State
	PoleTemp       map[types.ID]pole.State
	BulletTemp     map[types.ID]bullet.State
}

func NewWorld(s *Spawner, env types.Environment) *World {
	var game = &World{
		spawner:     s,
		environment: env,
		Drifters:    map[types.ID]*drifter.Drifter{},
		Dimetrodons: map[types.ID]*dimetrodon.Dimetrodon{},
		Poles:       map[types.ID]*pole.Pole{},
		Bullets:     map[types.ID]*bullet.Bullet{},

		DrifterTemp:    map[types.ID]drifter.State{},
		DimetrodonTemp: map[types.ID]dimetrodon.State{},
		PoleTemp:       map[types.ID]pole.State{},
		BulletTemp:     map[types.ID]bullet.State{},
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
		var newState = &drifter.State{}

		err := decoder.Decode(newState)
		if err != nil {
			return err
		}

		if exDrifter, ok := w.Drifters[id]; !ok { // new, so add it
			w.DrifterTemp[id] = *newState
		} else { // world
			exDrifter.State = *newState
		}
		break
	case dimetrodon.Class:
		var newState = &dimetrodon.State{}

		err := decoder.Decode(newState)
		if err != nil {
			return err
		}

		if exDimetrodon, ok := w.Dimetrodons[id]; !ok { // new, so add it
			w.DimetrodonTemp[id] = *newState
		} else { // update
			exDimetrodon.State = *newState
		}
		break
	case pole.Class:
		var newState = &pole.State{}

		err := decoder.Decode(newState)
		if err != nil {
			return err
		}

		if exPole, ok := w.Poles[id]; !ok {
			w.PoleTemp[id] = *newState
		} else { // update
			exPole.State = *newState
		}
		break
	case bullet.Class:
		var newState = &bullet.State{}

		err := decoder.Decode(newState)
		if err != nil {
			return err
		}

		if exBullet, ok := w.Bullets[id]; !ok {
			w.BulletTemp[id] = *newState
		} else { // update
			exBullet.State = *newState
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
