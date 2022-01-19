package ws

import (
	"encoding/gob"
)

type System interface {
	GetChannel() string
	GetData(*gob.Encoder) error
}

func (m *Manager) AddSystem(s *System) {
	var channel = (*s).GetChannel()
	m.systems[channel] = s
	m.AddEncoder(channel)
}

func (m *Manager) RemoveSystem(channel string) {
	delete(m.systems, channel)
}
