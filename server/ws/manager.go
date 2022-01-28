package ws

import (
	"github.com/josephnormandev/murder/common/communications"
	"github.com/josephnormandev/murder/common/types"
	"sync"
)

type Manager struct {
	timestamp types.Tick
	codecs    map[string]*communications.Codec

	systems   map[string]*System
	listeners map[string]*Listener
	spawns    map[int]*Spawn
	classes   map[string]int

	systemMutex sync.Mutex
	spawnMutex  sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		timestamp: 0,
		codecs:    map[string]*communications.Codec{},

		systems:   map[string]*System{},
		listeners: map[string]*Listener{},
		spawns:    map[int]*Spawn{},
		classes:   map[string]int{},

		systemMutex: sync.Mutex{},
		spawnMutex:  sync.Mutex{},
	}
}
