package ws

import "encoding/gob"

type System interface {
	GetID() int
	GetChannel() string
	GetData(*gob.Encoder) error
}

func (m *Manager) AddSystem(s *System) {
	var channel = (*s).GetChannel()
	m.systems[channel] = s

	for _, codec := range m.codecs {
		codec.AddEncoder(channel)
	}
}

func (m *Manager) RemoveSystem(channel string) {
	delete(m.systems, channel)
}
