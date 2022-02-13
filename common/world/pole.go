package world

import (
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

func (w *World) AddPole(p *pole.Pole) types.ID {
	var id = (*w.spawner).SpawnPole(p)
	w.Poles[id] = p
	return id
}

func (w *World) RemovePole(id types.ID) {
	(*w.spawner).DespawnPole(id)
	w.deletions.DeleteID(id)
	delete(w.Poles, id)
}
