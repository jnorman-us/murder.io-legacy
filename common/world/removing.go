package world

import "github.com/josephnormandev/murder/common/types"

func (w *World) RemoveDimetrodon(id types.ID) {
	(*w.spawner).DespawnDimetrodon(id)
	delete(w.Dimetrodons, id)
}

func (w *World) RemovePole(id types.ID) {
	(*w.spawner).DespawnPole(id)
	delete(w.Poles, id)
}

func (w *World) RemoveBullet(id types.ID) {
	(*w.spawner).DespawnBullet(id)
	delete(w.Bullets, id)
}
