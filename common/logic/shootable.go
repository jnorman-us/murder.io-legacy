package logic

import "github.com/josephnormandev/murder/common/types"

// Shootable is a shared logic interface, allowing players to fire their
// weapons and abilities
type Shootable interface {
	GetInput() types.Input
	GetShootingCoolDown() *types.CoolDown
	Shoot()
}

func (m *Manager) AddShootable(id types.ID, s *Shootable) {
	m.Shootables[id] = s
}

func (m *Manager) RemoveShootable(id types.ID) {
	delete(m.Shootables, id)
}

func (m *Manager) ShootingLogic(s *Shootable) {
	var shootable = *s
	var inputs = shootable.GetInput()
	var coolDown = shootable.GetShootingCoolDown()

	if inputs.AttackClick && coolDown.Ready() {
		shootable.Shoot()
		coolDown.Reset()
	}
	coolDown.Tick()
}
