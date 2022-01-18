package world

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/arrow"
	"github.com/josephnormandev/murder/common/entities/bow"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/entities/wall"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/server/websocket"
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

	network    *websocket.Manager
	collisions *collisions.Manager
	engine     *engine.Engine
	logic      *logic.Manager
}

func NewWorld(e *engine.Engine, l *logic.Manager, c *collisions.Manager, n *websocket.Manager) *World {
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
