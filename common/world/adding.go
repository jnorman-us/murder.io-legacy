package world

import (
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

func (w *World) AddDrifter(d *drifter.Drifter) types.ID {
	d.MarkDirty()
	var id = (*w.spawner).SpawnDrifter(d)
	w.Drifters[id] = d
	if w.environment.IsServer() {
		w.additions.Add(id, drifter.Class)
	}
	return id
}

func (w *World) AddDimetrodon(d *dimetrodon.Dimetrodon) types.ID {
	d.MarkDirty()
	var id = (*w.spawner).SpawnDimetrodon(d)
	w.Dimetrodons[id] = d
	if w.environment.IsServer() {
		w.additions.Add(id, dimetrodon.Class)
	}
	return id
}

func (w *World) AddPole(p *pole.Pole) types.ID {
	var id = (*w.spawner).SpawnPole(p)
	w.Poles[id] = p
	if w.environment.IsServer() {
		w.additions.Add(id, pole.Class)
	}
	return id
}

func (w *World) AddBullet(b *bullet.Bullet) types.ID {
	var id = (*w.spawner).SpawnBullet(b)
	w.Bullets[id] = b
	if w.environment.IsServer() {
		w.additions.Add(id, bullet.Class)
	}
	return id
}
