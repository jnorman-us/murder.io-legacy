package tcp

import "github.com/josephnormandev/murder/common/packet"

type Manager struct {
	packet.Codec

	systems   map[string]*System
	listeners map[string]*Listener
	spawns    map[int]*Spawn
}

func NewManager() *Manager {
	var codec = packet.NewCodec()
	return &Manager{
		Codec: *codec,

		systems:   map[string]*System{},
		listeners: map[string]*Listener{},
		spawns:    map[int]*Spawn{},
	}
}
