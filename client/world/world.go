package world

import (
	"encoding/gob"
	"fmt"
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
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
	input  *input.Manager
}

func NewWorld(e *engine.Engine, d *drawer.Drawer, i *input.Manager) *World {
	return &World{
		environment: types.ClientEnvironment(),
		tick:        0,
		Walls:       map[int]*wall.Wall{},
		Innocents:   map[int]*innocent.Innocent{},
		Swords:      map[int]*sword.Sword{},
		Bows:        map[int]*bow.Bow{},
		Arrows:      map[int]*arrow.Arrow{},

		input:  i,
		drawer: d,
		engine: e,
	}
}

func (w *World) GetClasses() []string {
	return []string{"innocent", "arrow", "wall", "sword", "bow"}
}

func (w *World) HandleSpawn(id int, class string, decoder *gob.Decoder) error {
	fmt.Println("HandleSpawn!", id, class)
	switch class {
	case "innocent":
		var newInn = &innocent.Innocent{}
		err := decoder.Decode(newInn)
		if err != nil {
			return err
		}

		var existing, ok = w.Innocents[id]
		if !ok { // add them
			w.AddInnocent(newInn)
			fmt.Printf("Adding %v\n", newInn)
		} else {
			fmt.Printf("Updating %v\n", existing)
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
	return nil
}
