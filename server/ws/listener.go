package ws

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/types"
)

type Listener interface {
	GetChannel() byte
	HandleData(types.UserID, *gob.Decoder) error // client, decoder
}

func (l *Lobby) AddListener(li *Listener) {
	var channel = (*li).GetChannel()
	l.listeners[channel] = li

	for _, c := range l.clients {
		c.codec.AddDecoder(channel)
	}
}

func (l *Lobby) RemoveListener(channel byte) {
	delete(l.listeners, channel)
}
