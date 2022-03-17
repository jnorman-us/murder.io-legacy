package ws

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/types"
)

type System interface {
	GetChannel() types.Channel
	GetData() []data.Data
}

func (l *Lobby) AddSystem(s *System) {
	l.Lock()
	defer l.Unlock()

	var channel = (*s).GetChannel()
	l.systems[channel] = s
}

func (l *Lobby) RemoveSystem(channel types.Channel) {
	l.Lock()
	defer l.Unlock()

	delete(l.systems, channel)
}
