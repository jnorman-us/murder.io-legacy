package dimetrodon

import "github.com/josephnormandev/murder/common/types"

type Spawner interface {
	DimetrodonShootBullet(*Dimetrodon)
	RemoveDimetrodon(types.ID)
}

func (d *Dimetrodon) SetSpawner(s *Spawner) {
	d.spawner = s
}
