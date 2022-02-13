package logic

import (
	"github.com/josephnormandev/murder/common/types"
)

type Manager struct {
	Tickables  map[types.ID]*Tickable
	Fireables  map[types.ID]*Fireable
	Shootables map[types.ID]*Shootable
	Driveables map[types.ID]*Driveable
}

func NewManager() *Manager {
	var manager = &Manager{
		Tickables:  map[types.ID]*Tickable{},
		Fireables:  map[types.ID]*Fireable{},
		Shootables: map[types.ID]*Shootable{},
		Driveables: map[types.ID]*Driveable{},
	}
	return manager
}

func (m *Manager) Tick() {
	for _, s := range m.Fireables {
		m.FireableLogic(s)
	}
	for _, s := range m.Shootables {
		m.ShootingLogic(s)
	}
	for _, d := range m.Driveables {
		m.Drive(d)
	}
	for _, t := range m.Tickables {
		(*t).Tick()
	}
}
