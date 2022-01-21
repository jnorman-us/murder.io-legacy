package ws

import "github.com/josephnormandev/murder/common/packet"

type Manager struct {
	codecs map[string]*packet.Codec

	systems   map[string]*System
	listeners map[string]*Listener
	spawns    map[int]*Spawn
	classes   map[string]int
}

func NewManager() *Manager {
	return &Manager{
		codecs: map[string]*packet.Codec{},

		systems:   map[string]*System{},
		listeners: map[string]*Listener{},
		spawns:    map[int]*Spawn{},
		classes:   map[string]int{},
	}
}
