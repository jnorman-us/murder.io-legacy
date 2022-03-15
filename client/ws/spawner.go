package ws

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/types"
)

type Spawner interface {
	HandleAddition(id types.ID, channel types.Channel, datum data.Data)
	HandleUpdate(id types.ID, channel types.Channel, datum data.Data)
	HandleDeletion(id types.ID, channel types.Channel, datum data.Data)
}

func (m *Manager) SetSpawner(s *Spawner) {
	m.spawner = s
}
