package ws

import "encoding/gob"

type System interface {
	GetChannel() string
	Flush() // for aggregating systems, bookmark data collection
	GetData(*gob.Encoder) error
}

func (m *Manager) AddSystem(s *System) {
	m.systemMutex.Lock()
	defer m.systemMutex.Unlock()

	var channel = (*s).GetChannel()
	m.systems[channel] = s

	for _, codec := range m.codecs {
		codec.AddEncoder(channel)
	}
}

func (m *Manager) RemoveSystem(channel string) {
	m.systemMutex.Lock()
	defer m.systemMutex.Unlock()

	delete(m.systems, channel)
}
