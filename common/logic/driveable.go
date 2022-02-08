package logic

import "github.com/josephnormandev/murder/common/types"

// Driveable is a shared logic interface, allowing the same driving
// logic to be applied to different driveable Cars
type Driveable interface {
	GetInput() types.Input
	ApplyForce(types.Vector)
	ApplyPositionalForce(types.Vector, types.Vector)
}

func (m *Manager) AddDriveable(id types.ID, d *Driveable) {
	m.Driveables[id] = d
}

func (m *Manager) RemoveDriveable(id types.ID) {
	delete(m.Driveables, id)
}

func (m *Manager) Drive(d *Driveable) {
	// TO DO, driving logic

	// parse the input

	// apply forces at nose, rear

	// consider drifting
}
