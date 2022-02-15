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

	systems   map[byte]*System
	listeners map[byte]*Listener
	spawns    map[types.ID]*Spawn
	classes   map[byte]int

	worldLock   *sync.Mutex
	systemMutex sync.Mutex
	spawnMutex  sync.Mutex
}

func NewLobby(info *LobbyInfo, lock *sync.Mutex) *Lobby {
	return &Lobby{
		info:      info,
		timestamp: 0,
		clients:   map[types.UserID]*Client{},

		systems:   map[byte]*System{},
		listeners: map[byte]*Listener{},
		spawns:    map[types.ID]*Spawn{},
		classes:   map[byte]int{},

		worldLock:   lock,
		systemMutex: sync.Mutex{},
		spawnMutex:  sync.Mutex{},
	}
}

func (l *Lobby) Send() {
	var ms = time.Duration(1000 / 20)
	for range time.Tick(ms * time.Millisecond) {
		l.worldLock.Lock()
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
		l.worldLock.Unlock()
		l.timestamp++
	}
}
