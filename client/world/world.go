package world

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/arrow"
	"github.com/josephnormandev/murder/common/entities/bow"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/entities/wall"
	"github.com/josephnormandev/murder/common/types"
)

type World struct {
	environment types.Environment
	tick        int

	Walls     map[int]*wall.Wall
	Innocents map[int]*innocent.Innocent
	Swords    map[int]*sword.Sword
	Bows      map[int]*bow.Bow
	Arrows    map[int]*arrow.Arrow

	drawer *drawer.Drawer
	engine *engine.Engine
}

func NewWorld(e *engine.Engine, d *drawer.Drawer) *World {
	return &World{
		environment: types.ClientEnvironment(),
		tick:        0,
		Walls:       map[int]*wall.Wall{},
		Innocents:   map[int]*innocent.Innocent{},
		Swords:      map[int]*sword.Sword{},
		Bows:        map[int]*bow.Bow{},
		Arrows:      map[int]*arrow.Arrow{},

		drawer: d,
		engine: e,
	}
}

func (w *World) GetClasses() []string {
	return []string{"innocent", "arrow", "wall", "sword", "bow"}
}

func (w *World) HandleSpawn(id int, class string, decoder *gob.Decoder) error {
	switch class {
	case "innocent":
		var newInn = &innocent.Innocent{}
		err := decoder.Decode(newInn)
		if err != nil {
			return err
		}

		var existing, ok = w.Innocents[id]
		if !ok { // add them
			newInn.Setup()
			w.AddInnocent(newInn)
		} else {
			existing.CopyKinetics(newInn.Collider)
		}
		break
	case "arrow":
		var newArr = &arrow.Arrow{}
		err := decoder.Decode(newArr)
		if err != nil {
			return err
		}

		var existing, ok = w.Arrows[id]
		if !ok {
			newArr.Setup()
			w.AddArrow(newArr)
		} else {
			existing.CopyKinetics(newArr.Collider)
		}
		break
	case "wall":
		var newWall = &wall.Wall{}
		err := decoder.Decode(newWall)
		if err != nil {
			return err
		}

		var existing, ok = w.Walls[id]
		if !ok {
			newWall.Setup()
			w.AddWall(newWall)
		} else {
			existing.CopyKinetics(newWall.Collider)
		}
		break
	case "sword":
		var newSword = &sword.Sword{}
		err := decoder.Decode(newSword)
		if err != nil {
			return err
		}

		var existing, ok = w.Swords[id]
		if !ok {
			newSword.Setup()
			w.AddSword(newSword)
		} else {
			existing.CopyKinetics(newSword.Collider)
		}
		break
	case "bow":
		var newBow = &bow.Bow{}
		err := decoder.Decode(newBow)
		if err != nil {
			return err
		}

		var existing, ok = w.Bows[id]
		if !ok {
			newBow.Setup()
			w.AddBow(newBow)
		} else {
			existing.CopyKinetics(newBow.Collider)
		}
		break
	}
	return nil
}

func (w *World) GetChannel() string {
	return "dlt"
}

func (w *World) HandleData(decoder *gob.Decoder) error {
	var removedIDs = &map[int]int{}
	err := decoder.Decode(removedIDs)

	for id := range *removedIDs {
		if _, ok := w.Innocents[id]; ok {
			w.RemoveInnocent(id)
		} else if _, ok := w.Arrows[id]; ok {
			w.RemoveArrow(id)
		} else if _, ok := w.Walls[id]; ok {
			w.RemoveWall(id)
		} else if _, ok := w.Swords[id]; ok {
			w.RemoveSword(id)
		} else if _, ok := w.Bows[id]; ok {
			w.RemoveBow(id)
		}
	}

	if err != nil {
		return err
	}
	return nil
}

func (w *World) GetCenterable(username string) *drawer.Centerable {
	for _, i := range w.Innocents {
		var inn = *i
		if inn.Username == username {
			var centerable = drawer.Centerable(i)
			return &centerable
		}
	}
	return nil
}
