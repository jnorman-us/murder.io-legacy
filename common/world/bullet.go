package world

import (
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/types"
)

func (w *World) AddBullet(b *bullet.Bullet) types.ID {
	var id = (*w.spawner).SpawnBullet(b)
	w.Bullets[id] = b
	return id
}

func (w *World) RemoveBullet(id types.ID) {
	(*w.spawner).DespawnBullet(id)
	w.deletions.DeleteID(id)
	delete(w.Bullets, id)
}
