package ws

import "encoding/gob"

type Spawn interface {
	GetID() int
	GetClass() string
	GetData(*gob.Encoder) error
}

func (m *Lobby) AddSpawn(id int, s *Spawn) {
	m.spawnMutex.Lock()
	defer m.spawnMutex.Unlock()

	var class = (*s).GetClass()
	var _, ok = m.classes[class]

	m.spawns[id] = s
	m.classes[class] = 0

	if !ok {
		for _, c := range m.clients {
			c.codec.AddEncoder(class)
		}
	}
}

func (m *Lobby) RemoveSpawn(id int) {
	m.spawnMutex.Lock()
	defer m.spawnMutex.Unlock()

	delete(m.spawns, id)
}
