package ws

import "encoding/gob"

type System interface {
	GetChannel() string
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

func (m *Lobby) RemoveSystem(channel string) {
	m.systemMutex.Lock()
	defer m.systemMutex.Unlock()

	delete(m.systems, channel)
}
