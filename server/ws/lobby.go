package ws

import (
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

func NewLobby(info *LobbyInfo) *Lobby {
	return &Lobby{
		Manager: *packets.NewManager(),
		info:    info,
		clients: map[types.UserID]*Client{},
	}
}

func (l *Lobby) Send() {
	var clump = l.MarshalPackets()
	for _, c := range l.clients {
		if (*c).Active() {
			c.Send(clump)
		}
	}
	l.TimeTick()
}

func (l *Lobby) Receive(client *Client, clump packets.Clump) {

}
