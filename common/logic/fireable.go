package logic

import (
	"github.com/josephnormandev/murder/common/types"
)

// Dissolvable is a shared logic interface, allowing for bullets to drop
// off as they move, and other logic specific to damaging projectiles
type Dissolvable interface {
	Break()
	GetRange() float64
	GetInitialPosition() types.Vector
	GetPosition() types.Vector
}

func (m *Manager) AddDissolvable(id types.ID, d *Dissolvable) {
	m.Dissolvables[id] = d
}

func (m *Manager) RemoveDissolvable(id types.ID) {
	delete(m.Dissolvables, id)
}

func (m *Manager) Dissolve(d *Dissolvable) {
	var dissolvable = *d
	var initPos = dissolvable.GetInitialPosition()
	var currentPos = dissolvable.GetPosition()
	var distance = initPos.Distance(currentPos)

	if distance >= dissolvable.GetRange() {
		dissolvable.Break()
	}
}
