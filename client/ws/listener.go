package ws

import (
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/types/action"
)

type Handler interface {
	HandleAdd(types.Channel, packets.Data)
	HandleUpdate(types.Channel, packets.Data)
	HandleDelete(types.Channel, packets.Data)
}

type Listener struct {
	handler *Handler
	channel types.Channel
	data    map[types.ID]*packets.Data
}

func newListener(c types.Channel, h *Handler) *Listener {
	var listener = &Listener{
		handler: h,
		channel: c,
		data:    map[types.ID]*packets.Data{},
	}
	return listener
}

func (l *Listener) receive(packet packets.Packet, elapsed byte) {
	var id = packet.ID
	var channel = packet.Channel
	var floatDiff = packet.FloatDiff
	var intDiff = packet.IntDiff
	var stringDiff = packet.StringDiff

	var handler = *l.handler
	switch packet.Action {
	case action.Actions.Add:
		var data = packets.NewData(id, packet.Channel, nil)
		data.SetDiffs(floatDiff, intDiff, stringDiff)
		data.Trickle(elapsed + 1)
		handler.HandleAdd(channel, *data)
		l.data[id] = data
	case action.Actions.Update:
		if data, ok := l.data[id]; ok {
			data.SetDiffs(floatDiff, intDiff, stringDiff)
		}
	case action.Actions.Delete:
		if data, ok := l.data[id]; ok {
			data.SetDiffs(floatDiff, intDiff, stringDiff)
			data.Trickle(elapsed + 1)
			handler.HandleDelete(channel, *data)
		}
		delete(l.data, id)
	}
}

func (l *Listener) trickle(elapsed byte) {
	for _, data := range l.data {
		data.Trickle(elapsed)
		data.Print()
	}
}
