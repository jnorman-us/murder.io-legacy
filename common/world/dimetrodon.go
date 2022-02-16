package world

import (
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/types"
)

func (w *World) AddDimetrodon(d *dimetrodon.Dimetrodon) types.ID {
	var id = (*w.spawner).SpawnDimetrodon(d)
	w.Dimetrodons[id] = d
	return id
}

func (w *World) RemoveDimetrodon(id types.ID) {
	(*w.spawner).DespawnDimetrodon(id)
	w.deletions.DeleteID(id)
	delete(w.Dimetrodons, id)
}

func (w *World) DimetrodonShootBullet(d *dimetrodon.Dimetrodon) {
	var shooter = bullet.Shooter(d)
	var newBullet = bullet.NewBullet(&shooter, 0)
	w.AddBullet(newBullet)
}
