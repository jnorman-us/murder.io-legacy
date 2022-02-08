package game

import (
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/types"
)

func (g *Game) AddDrifter(id types.ID, d *drifter.Drifter) {
	d.ID = id
	g.drifters[id] = d
}

func (g *Game) RemoveDrifter(id types.ID) {
	delete(g.drifters, id)
}
