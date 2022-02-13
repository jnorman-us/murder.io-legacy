package world

import (
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/types"
)

func (w *World) AddDrifter(d *drifter.Drifter) types.ID {
	var id = (*w.spawner).SpawnDrifter(d)
	w.Drifters[id] = d
	return id
}

func (w *World) RemoveDrifter(id types.ID) {
	(*w.spawner).DespawnDrifter(id)
	w.deletions.DeleteID(id)
	delete(w.Drifters, id)
}

func (w *World) DrifterShootBullet(d *drifter.Drifter, angle float64) {
	var shooter = bullet.Shooter(d)
	var newBullet = bullet.NewBullet(&shooter, angle)
	w.AddBullet(newBullet)
}
