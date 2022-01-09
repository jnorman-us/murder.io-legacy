package packet

import "encoding/gob"

type Listener interface {
	GetChannel() string
	HandleData(*gob.Decoder)
}

func (m *Manager) AddListener(channel string, l *Listener) int {
	m.Listeners[channel] = append(m.Listeners[channel], l)
	return len(m.Listeners[channel]) - 1
}

func (m *Manager) RemoveListener(channel string, index int) {
	m.Listeners[channel] = append(m.Listeners[channel][:index], m.Listeners[channel][index+1:]...)
}
