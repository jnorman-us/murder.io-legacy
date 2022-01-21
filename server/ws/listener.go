package ws

import "encoding/gob"

type Listener interface {
	GetChannel() string
	HandleData(string, *gob.Decoder) error // client, decoder
}

func (m *Manager) AddListener(l *Listener) {
	var channel = (*l).GetChannel()
	m.listeners[channel] = l

	for _, codec := range m.codecs {
		codec.AddDecoder(channel)
	}
}

func (m *Manager) RemoveListener(channel string) {
	delete(m.listeners, channel)
}
