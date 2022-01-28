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
	"github.com/josephnormandev/murder/server/input"
	"github.com/josephnormandev/murder/server/world/deletions"
	"github.com/josephnormandev/murder/server/ws"
)

type World struct {
	environment types.Environment
	currentID   int
	Walls       map[int]*wall.Wall
	Innocents   map[int]*innocent.Innocent
	Swords      map[int]*sword.Sword
	Bows        map[int]*bow.Bow
	Arrows      map[int]*arrow.Arrow

	network    *ws.Manager
	collisions *collisions.Manager
	Deletions  *deletions.Manager
	engine     *engine.Engine
	logic      *logic.Manager
	inputs     *input.Manager
}

func NewWorld(e *engine.Engine, l *logic.Manager, c *collisions.Manager, n *ws.Manager, i *input.Manager) *World {
	return &World{
		environment: types.ServerEnvironment(),
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
		inputs:     i,

		Deletions: deletions.NewManager(),
	}
}

func (w *World) NextAvailableID() int {
	w.currentID++
	return w.currentID
}

func (w *World) ResetInnocents() {
	var toRemove []int

	for id := range w.Innocents {
		toRemove = append(toRemove, id)
	}

	for _, id := range toRemove {
		w.RemoveInnocent(id)
	}
}
