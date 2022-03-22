package ws

import (
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
)

type Listener interface {
	GetChannel() types.Channel
	HandleData([]packets.Data) // id, decoder
}

type FutureListener interface {
	GetChannel() types.Channel
	HandleFutureData([]packets.Data) // decoder, ttl
}

func (m *Manager) AddListener(l *Listener) {
	var channel = (*l).GetChannel()
	m.listeners[channel] = l
}

func (m *Manager) RemoveListener(channel types.Channel) {
	delete(m.listeners, channel)
}

func (m *Manager) AddFutureListener(l *FutureListener) {
	var channel = (*l).GetChannel()
	m.futureListeners[channel] = l
}

func (m *Manager) RemoveFutureListener(channel types.Channel) {
	delete(m.futureListeners, channel)
}
