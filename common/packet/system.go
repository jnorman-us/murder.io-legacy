package packet

import (
	"bytes"
	"encoding/gob"
)

type System interface {
	GetChannel() string
	GetData(*gob.Encoder)
}

func (m *Manager) AddSystem(channel string, s *System) {
	m.Systems[channel] = s

	if _, ok := m.outputs[channel]; !ok {
		var channelOutput = new(bytes.Buffer)
		m.outputs[channel] = channelOutput
		m.encoders[channel] = gob.NewEncoder(channelOutput)
	}
}

func (m *Manager) RemoveSystem(channel string) {
	delete(m.Systems, channel)
}
