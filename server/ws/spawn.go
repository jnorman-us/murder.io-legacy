package ws

import "encoding/gob"

type Spawn interface {
	GetID() int
	GetClass() string
	GetData(*gob.Encoder) error
}

func (m *Manager) AddSpawn(id int, s *Spawn) {
	var class = (*s).GetClass()
	var _, ok = m.classes[class]

	m.spawns[id] = s
	m.classes[class] = 0

	if !ok {
		for _, codec := range m.codecs {
			codec.AddEncoder(class)
		}
	}
}

func (m *Manager) RemoveSpawn(id int) {
	delete(m.spawns, id)
}
