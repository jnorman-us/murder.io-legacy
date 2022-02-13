package ws

import "encoding/gob"

type System interface {
	GetChannel() byte
	Flush() // for aggregating systems, bookmark data collection
	GetData(*gob.Encoder) error
}

func (l *Lobby) AddSystem(s *System) {
	l.systemMutex.Lock()
	defer l.systemMutex.Unlock()

	var channel = (*s).GetChannel()
	l.systems[channel] = s

	for _, c := range l.clients {
		c.codec.AddEncoder(channel)
	}
}

func (l *Lobby) RemoveSystem(channel byte) {
	l.systemMutex.Lock()
	defer l.systemMutex.Unlock()

	delete(l.systems, channel)
}
