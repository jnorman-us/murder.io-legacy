package collisions

import "github.com/josephnormandev/murder/common/types"

type Manager struct {
	StaticDynamics   map[types.ID]*StaticDynamic
	DynamicStatics   map[types.ID]*DynamicStatic
	BulletShootables map[types.ID]*BulletShootable
	ShootableBullets map[types.ID]*ShootableBullet
}

func NewManager() *Manager {
	return &Manager{
		StaticDynamics:   map[types.ID]*StaticDynamic{},
		DynamicStatics:   map[types.ID]*DynamicStatic{},
		BulletShootables: map[types.ID]*BulletShootable{},
		ShootableBullets: map[types.ID]*ShootableBullet{},
	}
}

func (m *Manager) ResolveCollisions() {
	m.resolveStaticDynamics()
	m.resolveBulletShootables()
}
