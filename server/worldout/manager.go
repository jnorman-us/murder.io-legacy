package worldout

import (
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
)

type Manager struct {
	packets     *packets.Manager
	datafiables map[types.ID]*Datafiable
	streams     map[types.Channel]*packets.Stream
	data        map[types.ID]*packets.Data
}

func NewManager(pm *packets.Manager) *Manager {
	var manager = &Manager{
		packets:     pm,
		datafiables: map[types.ID]*Datafiable{},
		streams:     map[types.Channel]*packets.Stream{},
		data:        map[types.ID]*packets.Data{},
	}
	return manager
}

func (m *Manager) PollData() {
	for id, d := range m.datafiables {
		var datafiable = *d
		var data = m.data[id]

		datafiable.PopulateData(data)
	}
}
