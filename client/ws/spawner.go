package ws

import (
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
)

type Spawner interface {
	HandleAddition(id types.ID, channel types.Channel, datum packets.Data)
	HandleUpdate(id types.ID, channel types.Channel, datum packets.Data)
	HandleDeletion(id types.ID, channel types.Channel, datum packets.Data)
}

func (m *Manager) SetSpawner(s *Spawner) {
	m.spawner = s
}
