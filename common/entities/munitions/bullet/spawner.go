package bullet

import "github.com/josephnormandev/murder/common/types"

type Spawner interface {
	RemoveBullet(types.ID)
	DisableBullet(types.ID)
}

func (b *Bullet) SetSpawner(s *Spawner) {
	b.spawner = s
}
