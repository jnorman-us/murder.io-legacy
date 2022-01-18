package tcp

import (
	"bufio"
	"github.com/josephnormandev/murder/common/packet"
	"net"
)

type Client struct {
	identifier string
	manager    *Manager
	socket     *net.TCPConn
	reader     *bufio.Reader
	writer     *bufio.Writer
}

func NewClient(identifier string, manager *Manager) *Client {
	return &Client{
		identifier: identifier,
		manager:    manager,
	}
}

func (c *Client) GetSocket() *net.TCPConn {
	return c.socket
}

func (c *Client) SetSocket(s *net.TCPConn) {
	c.socket = s
	c.reader = bufio.NewReader(s)
	c.writer = bufio.NewWriter(s)
}

func (c *Client) Connected() bool {
	return c.socket != nil
}

func (c *Client) EncodeSystems() {
	var manager = c.manager
	var packetArray []packet.Packet

	for _, s := range manager.systems {
		var system = *s

		var channel = system.GetChannel()
		var output = manager.Outputs[channel]
		var encoder = manager.Encoders[channel]

		output.Reset()
		system.GetData(encoder)
		var outputBytes = output.Bytes()

		packetArray = append(packetArray, packet.Packet{
			Channel: channel,
			Data:    outputBytes,
		})
	}

	for id, s := range manager.spawns {
		var spawn = *s

		var class = spawn.GetClass()
		var output = manager.Outputs[class]
		var encoder = manager.Encoders[class]

		output.Reset()
		spawn.GetData(encoder)
		var outputBytes = output.Bytes()

		packetArray = append(packetArray, packet.Packet{
			ID:      id,
			Channel: class,
			Data:    outputBytes,
		})
	}

	manager.EncodeOutputs(c.identifier, packetArray)
}

func (c *Client) DecodeForListeners(ps []packet.Packet) {
	var manager = c.manager

	for _, packet := range ps {
		var _ = packet.ID
		var channel = packet.Channel
		var data = packet.Data

		var l, ok = manager.listeners[channel]
		var input = manager.Inputs[channel]
		var decoder = manager.Decoders[channel]

		if ok {
			var listener = *l

			input.Reset()
			input.Write(data)
			listener.HandleData(c.identifier, decoder)
		}
	}
}
