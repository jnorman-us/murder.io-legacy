package ws

import (
	"github.com/josephnormandev/murder/common/packet"
	"sync"
)

type Manager struct {
	codecs map[string]*packet.Codec

	systems   map[string]*System
	listeners map[string]*Listener
	spawns    map[int]*Spawn
	classes   map[string]int

	systemMutex sync.Mutex
	spawnMutex  sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		codecs: map[string]*packet.Codec{},

		systems:   map[string]*System{},
		listeners: map[string]*Listener{},
		spawns:    map[int]*Spawn{},
		classes:   map[string]int{},

		systemMutex: sync.Mutex{},
		spawnMutex:  sync.Mutex{},
	}
}
