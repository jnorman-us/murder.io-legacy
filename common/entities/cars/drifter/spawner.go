package drifter

import "github.com/josephnormandev/murder/common/types"

type Spawner interface {
	DrifterShootBullet(*Drifter, float64)
	RemoveDrifter(types.ID)
}

func (d *Drifter) SetSpawner(s *Spawner) {
	d.spawner = s
}
