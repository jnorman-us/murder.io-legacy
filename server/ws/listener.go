package ws

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/types"
)

type Listener interface {
	GetChannel() types.Channel
	HandleAdd(id types.UserID, added data.Data)
	HandleUpdate(id types.UserID, updated data.Data)
	HandleDelete(id types.UserID, deleted data.Data)
}

func (l *Lobby) AddListener(li *Listener) {
	l.Lock()
	defer l.Unlock()

	var channel = (*li).GetChannel()
	l.listeners[channel] = li
}

func (l *Lobby) RemoveListener(channel types.Channel) {
	l.Lock()
	defer l.Unlock()

	delete(l.listeners, channel)
}
