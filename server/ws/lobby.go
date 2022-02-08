package ws

import (
	"fmt"
	"github.com/josephnormandev/murder/common/types"
	"sync"
	"time"
)

type Lobby struct {
	info      *LobbyInfo
	timestamp types.Tick
	clients   map[types.UserID]*Client // per user Client

	systems   map[string]*System
	listeners map[string]*Listener
	spawns    map[int]*Spawn
	classes   map[string]int

	systemMutex sync.Mutex
	spawnMutex  sync.Mutex
}

func NewLobby(info *LobbyInfo) *Lobby {
	return &Lobby{
		info:      info,
		timestamp: 0,
		clients:   map[types.UserID]*Client{},

		systems:   map[string]*System{},
		listeners: map[string]*Listener{},
		spawns:    map[int]*Spawn{},
		classes:   map[string]int{},

		systemMutex: sync.Mutex{},
		spawnMutex:  sync.Mutex{},
	}
}

func (l *Lobby) Send() {
	for range time.Tick(50 * time.Millisecond) {
		for _, s := range l.systems {
			var system = *s
			system.Flush()
		}
		for _, c := range l.clients {
			var client = *c
			if client.Active() {
				var packetCollection, err = client.EncodeSystems()
				if err != nil {
					fmt.Printf("Error with sending! %v\n", err)
				}
				client.Send(packetCollection)
			}
		}
	}
	l.timestamp++
}
