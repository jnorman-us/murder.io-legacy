package ws

import (
	"github.com/josephnormandev/murder/common/communications"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/types/action"
	"github.com/josephnormandev/murder/common/types/timestamp"
	"sync"
	"time"
)

type Lobby struct {
	sync.Mutex
	info    *LobbyInfo
	clients map[types.UserID]*Client // per user Client

	time      *timestamp.Timestamp
	systems   map[types.Channel]*System
	listeners map[types.Channel]*Listener

	spawns      map[types.ID]*Spawn
	additions   map[types.ID]*Spawn
	addTimes    map[types.ID]time.Duration
	deletions   map[types.ID]*Spawn
	deleteTimes map[types.ID]time.Duration
}

func NewLobby(info *LobbyInfo) *Lobby {
	return &Lobby{
		info:    info,
		time:    timestamp.NewTimestamp(),
		clients: map[types.UserID]*Client{},

		systems:   map[types.Channel]*System{},
		listeners: map[types.Channel]*Listener{},

		spawns:      map[types.ID]*Spawn{},
		additions:   map[types.ID]*Spawn{},
		addTimes:    map[types.ID]time.Duration{},
		deletions:   map[types.ID]*Spawn{},
		deleteTimes: map[types.ID]time.Duration{},
	}
}

func (l *Lobby) Send() {
	for _, c := range l.clients {
		if (*c).Active() {
			var clump = l.EncodeSystems(c)
			(*c).Send(clump)
		}
	}
	l.time.TimeTick()
}

// EncodeSystems allows the use of a per-client visibility filter
// so that each client receives a stream of bytes unique to itself
func (l *Lobby) EncodeSystems(c *Client) communications.Clump {
	l.Lock()
	defer l.Unlock()

	var packets []communications.Packet
	for id, s := range l.additions {
		var spawn = *s
		var packet = communications.NewSpawnPacket(
			spawn.GetID(),
			spawn.GetClass(),
			action.Actions.Add,
			l.addTimes[id],
			spawn.GetData(),
		)
		packets = append(packets, packet)
	}
	for id, s := range l.deletions {
		var spawn = *s
		var packet = communications.NewSpawnPacket(
			spawn.GetID(),
			spawn.GetClass(),
			action.Actions.Delete,
			l.deleteTimes[id],
			spawn.GetData(),
		)
		packets = append(packets, packet)
	}
	for _, s := range l.spawns {
		var spawn = *s
		var packet = communications.NewSpawnPacket(
			spawn.GetID(),
			spawn.GetClass(),
			action.Actions.Update,
			0.0,
			spawn.GetData(),
		)
		packets = append(packets, packet)
	}
	l.additions = map[types.ID]*Spawn{}
	l.addTimes = map[types.ID]time.Duration{}
	l.deletions = map[types.ID]*Spawn{}
	l.deleteTimes = map[types.ID]time.Duration{}

	return communications.Clump{
		Packets:   packets,
		Timestamp: l.time.Tick,
	}
}

func (l *Lobby) EncodeCatchupSystems(c *Client) communications.Clump {
	l.Lock()
	defer l.Unlock()

	var packets []communications.Packet
	for _, s := range l.spawns {
		var spawn = *s
		var packet = communications.NewSpawnPacket(
			spawn.GetID(),
			spawn.GetClass(),
			action.Actions.Add,
			0.0,
			spawn.GetData(),
		)
		packets = append(packets, packet)
	}

	return communications.Clump{
		Packets:   packets,
		Timestamp: l.time.Tick,
	}
}

func (l *Lobby) DecodeForListeners(c *Client, clump communications.Clump) error {
	/*var lobby = c.lobby
	var codec = c.codec

	for _, p := range clump.Packets {
		var channel = p.Channel
		var data = p.Data

		var l, ok = lobby.listeners[channel]
		if ok {
			var listener = *l
			decoder, err := codec.BeginDecode(channel, data)
			if err != nil {
				return err
			}
			err = listener.HandleData(c.identifier, decoder)
			if err != nil {
				return err
			}
			codec.EndDecode(channel)
		}
	}*/
	return nil
}
