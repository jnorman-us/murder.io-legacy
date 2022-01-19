package ws

import "encoding/gob"

type Spawner interface {
	HandleSpawn(int, string, *gob.Decoder) // id, class, decoder
}

func (m *Manager) SetSpawner(s *Spawner) {
	m.spawner = s
}
