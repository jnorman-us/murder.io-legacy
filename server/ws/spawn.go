package ws

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/types"
)

type Spawn interface {
	GetID() types.ID
	GetClass() byte
	GetData(*gob.Encoder) error
}

func (l *Lobby) AddSpawn(id types.ID, s *Spawn) {
	l.spawnMutex.Lock()
	defer l.spawnMutex.Unlock()

	var class = (*s).GetClass()
	var _, ok = l.classes[class]

	l.spawns[id] = s
	l.classes[class] = 0

	if !ok {
		for _, c := range l.clients {
			c.codec.AddEncoder(class)
		}
	}
}

func (l *Lobby) RemoveSpawn(id types.ID) {
	l.spawnMutex.Lock()
	defer l.spawnMutex.Unlock()

	delete(l.spawns, id)
}
