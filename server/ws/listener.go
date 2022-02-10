package ws

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/types"
)

type Listener interface {
	GetChannel() byte
	HandleData(types.UserID, *gob.Decoder) error // client, decoder
}

func (m *Lobby) AddListener(l *Listener) {
	var channel = (*l).GetChannel()
	m.listeners[channel] = l

	for _, c := range m.clients {
		c.codec.AddDecoder(channel)
	}
}

func (m *Lobby) RemoveListener(channel byte) {
	delete(m.listeners, channel)
}
