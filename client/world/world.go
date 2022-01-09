package world

import (
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

	input *input.Manager
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
