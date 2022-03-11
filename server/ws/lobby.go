package ws

import (
	"fmt"
	"github.com/josephnormandev/murder/common/types"
	"sync"
)

type Lobby struct {
	info    *LobbyInfo
	clients map[types.UserID]*Client // per user Client

	time      *types.Time
	systems   map[byte]*System
	listeners map[byte]*Listener
	spawns    map[types.ID]*Spawn
	classes   map[byte]int

	systemMutex sync.Mutex
	spawnMutex  sync.Mutex
}

func NewLobby(info *LobbyInfo, time *types.Time) *Lobby {
	return &Lobby{
		info:    info,
		time:    time,
		clients: map[types.UserID]*Client{},

		systems:   map[byte]*System{},
		listeners: map[byte]*Listener{},
		spawns:    map[types.ID]*Spawn{},
		classes:   map[byte]int{},

		systemMutex: sync.Mutex{},
		spawnMutex:  sync.Mutex{},
	}
}

func (l *Lobby) Send() {
	for _, s := range l.systems {
		var system = *s
		system.Flush()
	}
	for _, c := range l.clients {
		if (*c).Active() {
			var packetCollection, err = (*c).EncodeSystems()
			if err != nil {
				fmt.Printf("Error with sending! %v\n", err)
			}
			(*c).Send(packetCollection)
		}
	}
}
