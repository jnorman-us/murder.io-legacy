package ws

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/types"
)

type Listener interface {
	GetChannel() types.Channel
	HandleData([]data.Data) // id, decoder
}

type FutureListener interface {
	GetChannel() types.Channel
	HandleFutureData([]data.Data) // decoder, ttl
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
