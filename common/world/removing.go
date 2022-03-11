package world

import "github.com/josephnormandev/murder/common/types"

func (w *World) RemoveDrifter(id types.ID) {
	(*w.spawner).DespawnDrifter(id)
	if w.environment.IsServer() {
		w.deletions.Delete(id)
	}
	delete(w.Drifters, id)
}

func (w *World) RemoveDimetrodon(id types.ID) {
	(*w.spawner).DespawnDimetrodon(id)
	if w.environment.IsServer() {
		w.deletions.Delete(id)
	}
	delete(w.Dimetrodons, id)
}

func (w *World) RemovePole(id types.ID) {
	(*w.spawner).DespawnPole(id)
	if w.environment.IsServer() {
		w.deletions.Delete(id)
	}
	delete(w.Poles, id)
}

func (w *World) RemoveBullet(id types.ID) {
	(*w.spawner).DespawnBullet(id)
	if w.environment.IsServer() {
		w.deletions.Delete(id)
	}
	delete(w.Bullets, id)
}
