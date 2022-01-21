package ws

import (
	"encoding/gob"
)

type Spawner interface {
	GetClasses() []string
	HandleSpawn(int, string, *gob.Decoder) error // id, class, decoder
}

func (m *Manager) SetSpawner(s *Spawner) {
	for _, class := range (*s).GetClasses() {
		m.AddDecoder(class)
	}
	m.spawner = s
}
