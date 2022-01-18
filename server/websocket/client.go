package websocket

import "github.com/josephnormandev/murder/common/packet"

type Client struct {
	identifier string
	manager    *Manager
}

func NewClient(id string, m *Manager) *Client {
	return &Client{
		identifier: id,
		manager:    m,
	}
}

// EncodeSystems allows the use of a per-client visibility filter
// so that each client receives a stream of bytes unique to itself
func (c *Client) EncodeSystems() {
	var packetArray []packet.Packet
	var manager = c.manager

	// perhaps implement this in the manager because
	// there is no need to make this data unique per client
	for _, s := range manager.systems {
		var system = *s
		var channel = system.GetChannel()

		var encoder = manager.BeginEncode(channel)
		system.GetData(encoder)
		var outputBytes = manager.EndEncode(channel)

		packetArray = append(packetArray, packet.Packet{
			ID:      -1,
			Channel: channel,
			Data:    outputBytes,
		})
	}

	for _, s := range manager.spawns {
		var spawn = *s
		var id = spawn.GetID()
		var class = spawn.GetClass()

		var encoder = manager.BeginEncode(class)
		spawn.GetData(encoder)
		var outputBytes = manager.EndEncode(class)

		packetArray = append(packetArray, packet.Packet{
			ID:      id,
			Channel: class,
			Data:    outputBytes,
		})
	}
}

func (c *Client) DecodeForListeners(ps []packet.Packet) {
	var manager = c.manager

	for _, packet := range ps {
		var channel = packet.Channel
		var data = packet.Data

		var l, ok = manager.listeners[channel]
		if ok {
			var decoder = manager.BeginDecode(channel, data)
			var listener = *l
			listener.HandleData(c.identifier, decoder)
			manager.EndDecode(channel)
		}
	}
}
