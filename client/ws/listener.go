package ws

import (
	"encoding/gob"
	"time"
)

type Listener interface {
	GetChannel() byte
	HandleData(*gob.Decoder) error // id, decoder
}

type FutureListener interface {
	GetChannel() byte
	HandleFutureData(*gob.Decoder, time.Duration) error // decoder, ttl
}

func (m *Manager) AddListener(l *Listener) {
	var channel = (*l).GetChannel()
	m.listeners[channel] = l
	m.AddDecoder(channel)
}

func (m *Manager) RemoveListener(channel byte) {
	delete(m.listeners, channel)
}

func (m *Manager) AddFutureListener(l *FutureListener) {
	var channel = (*l).GetChannel()
	m.futureListeners[channel] = l
	m.AddDecoder(channel)
}

func (m *Manager) RemoveFutureListener(channel byte) {
	delete(m.futureListeners, channel)
}
