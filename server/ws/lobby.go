package ws

import (
	"fmt"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"sync"
)

type Lobby struct {
	sync.Mutex
	packets.Manager
	info    *LobbyInfo
	clients map[types.UserID]*Client // per user Client
}

func NewLobby(info *LobbyInfo, packets *packets.Manager) *Lobby {
	return &Lobby{
		Manager: *packets,
		info:    info,
		clients: map[types.UserID]*Client{},
	}
}

func (l *Lobby) Send() {
	var clump = l.MarshalPackets()
	fmt.Println("Clump", clump)
	/*
		for _, c := range l.clients {
			if (*c).Active() {
				c.Send(clump)
			}
		}
		fmt.Println("REst")*/
}
