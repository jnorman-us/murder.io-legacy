package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/wall"
)

type World struct {
	tick      int
	currentID int
	Walls     map[int]*wall.Wall
	Innocents map[int]*innocent.Innocent

	drawer     *drawer.Drawer
	collisions *collisions.Manager
	engine     *engine.Engine

	input *input.Manager
}

func NewClientWorld(e *engine.Engine, c *collisions.Manager, d *drawer.Drawer, i *input.Manager) *World {
	return &World{
		tick:      0,
		currentID: 0,
		Walls:     map[int]*wall.Wall{},
		Innocents: map[int]*innocent.Innocent{},

		input:      i,
		drawer:     d,
		engine:     e,
		collisions: c,
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
