package packet

import "encoding/gob"

// Spawner is the interface that the Network Manager uses to
// update the world state. Listener type
type Spawner interface {
	AddSpawn(int, string, *gob.Decoder)
	UpdateSpawn(int, *gob.Decoder)
	RemoveSpawn(int)
}

func (m *Manager) SetSpawner(s *Spawner) {
	m.Spawner = s
}
