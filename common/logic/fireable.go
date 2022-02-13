package logic

import (
	"github.com/josephnormandev/murder/common/types"
)

// Fireable is a shared logic interface, allowing for bullets to drop
// off as they move, and other logic specific to damaging projectiles
type Fireable interface {
	Destroy()
	GetRange() float64
	GetInitialPosition() types.Vector
	GetPosition() types.Vector
}

func (m *Manager) AddFireable(id types.ID, s *Fireable) {
	m.Fireables[id] = s
}

func (m *Manager) RemoveFireable(id types.ID) {
	delete(m.Fireables, id)
}

func (m *Manager) FireableLogic(f *Fireable) {
	var fireable = *f
	var initPos = fireable.GetInitialPosition()
	var currentPos = fireable.GetPosition()
	var distance = initPos.Distance(currentPos)

	if distance >= fireable.GetRange() {
		fireable.Destroy()
	}
}
