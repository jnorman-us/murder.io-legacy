package tcp

import "encoding/gob"

type Spawn interface {
	GetID() int
	GetClass() string
	GetData(*gob.Encoder)
}

func (m *Manager) AddSpawn(id int, s *Spawn) {
	var class = (*s).GetClass()
	m.spawns[id] = s

	var _, ok = m.Encoders[class]

	if !ok {
		m.AddEncoder(class)
	}
}

func (m *Manager) RemoveSpawn(id int) {
	delete(m.spawns, id)
}
