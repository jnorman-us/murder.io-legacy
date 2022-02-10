package ws

import "encoding/gob"

type System interface {
	GetChannel() byte
	Flush() // for aggregating systems, bookmark data collection
	GetData(*gob.Encoder) error
}

func (m *Lobby) AddSystem(s *System) {
	m.systemMutex.Lock()
	defer m.systemMutex.Unlock()

	var channel = (*s).GetChannel()
	m.systems[channel] = s

	for _, c := range m.clients {
		c.codec.AddEncoder(channel)
	}
}

func (m *Lobby) RemoveSystem(channel byte) {
	m.systemMutex.Lock()
	defer m.systemMutex.Unlock()

	delete(m.systems, channel)
}
