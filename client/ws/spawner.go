package ws

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/types"
)

type Spawner interface {
	GetClasses() []byte
	HandleSpawn(types.ID, byte, *gob.Decoder) error // id, class, decoder
}

func (m *Manager) SetSpawner(s *Spawner) {
	for _, class := range (*s).GetClasses() {
		m.AddDecoder(class)
	}
	m.spawner = s
}
