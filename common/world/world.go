package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/arrow"
	"github.com/josephnormandev/murder/common/entities/bow"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/entities/wall"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/packet"
	"github.com/josephnormandev/murder/common/types"
)

type World struct {
	environment types.Environment
	tick        int
	currentID   int
	Walls       map[int]*wall.Wall
	Innocents   map[int]*innocent.Innocent
	Swords      map[int]*sword.Sword
	Bows        map[int]*bow.Bow
	Arrows      map[int]*arrow.Arrow

	network    *packet.Manager
	drawer     *drawer.Drawer
	collisions *collisions.Manager
	engine     *engine.Engine
	logic      *logic.Manager

	input *input.Manager
}

func NewClientWorld(e *engine.Engine, d *drawer.Drawer, i *input.Manager) *World {
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

func NewServerWorld(e *engine.Engine, l *logic.Manager, c *collisions.Manager, n *packet.Manager) *World {
	return &World{
		environment: types.ServerEnvironment(),
		tick:        0,
		currentID:   0,
		Walls:       map[int]*wall.Wall{},
		Innocents:   map[int]*innocent.Innocent{},
		Swords:      map[int]*sword.Sword{},
		Bows:        map[int]*bow.Bow{},
		Arrows:      map[int]*arrow.Arrow{},

		logic:      l,
		collisions: c,
		network:    n,
		engine:     e,
	}
}

func (w *World) NextAvailableID() int {
	w.currentID++
	return w.currentID
}

func (w *World) Tick() {
	w.tick++
}

func (w *World) GetTick() int {
	return w.tick
}
