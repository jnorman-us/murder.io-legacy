package ws

import (
	"encoding/gob"
	"time"
)

type Listener interface {
	GetChannel() string
	HandleData(*gob.Decoder) error // id, decoder
}

type FutureListener interface {
	GetChannel() string
	HandleFutureData(*gob.Decoder, time.Duration) error // decoder, ttl
}

func (m *Manager) AddListener(l *Listener) {
	var channel = (*l).GetChannel()
	m.listeners[channel] = l
	m.AddDecoder(channel)
}

func (m *Manager) RemoveListener(channel string) {
	delete(m.listeners, channel)
}

func (m *Manager) AddFutureListener(l *FutureListener) {
	var channel = (*l).GetChannel()
	m.futureListeners[channel] = l
	m.AddDecoder(channel)
}

func (m *Manager) RemoveFutureListener(channel string) {
	delete(m.futureListeners, channel)
}
