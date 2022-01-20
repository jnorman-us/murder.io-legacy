package ws

import (
	"context"
	"fmt"
	"github.com/josephnormandev/murder/common/packet"
	"golang.org/x/sync/errgroup"
	"nhooyr.io/websocket"
	"time"
)

type Client struct {
	identifier string
	manager    *Manager

	cancel    func()
	connected bool

	packetsOut chan []packet.Packet
}

func NewClient(id string, m *Manager) *Client {
	return &Client{
		identifier: id,
		manager:    m,
		packetsOut: make(chan []packet.Packet),
	}
}

func (c *Client) Setup(ctx context.Context, conn *websocket.Conn) error {
	ctx2, cancel := context.WithCancel(ctx)
	c.connected = true
	c.cancel = cancel
	defer c.Close()

	group, grpCtx := errgroup.WithContext(ctx2)

	group.Go(func() error {
		return c.Write(grpCtx, conn)
	})
	group.Go(func() error {
		return c.Read(grpCtx, conn)
	})

	err := group.Wait()
	return err
}

func (c *Client) Write(parentCtx context.Context, conn *websocket.Conn) error {
	var manager = c.manager
	for {
		select {
		case packets := <-c.packetsOut:
			byteArray, err := manager.EncodeOutputs(packets)
			if err != nil {
				return err
			}
			err = conn.Write(parentCtx, websocket.MessageBinary, byteArray)
			if err != nil {
				return err
			}
		case <-parentCtx.Done():
			return parentCtx.Err()
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (c *Client) Read(parentCtx context.Context, conn *websocket.Conn) error {
	var manager = c.manager

	for {
		select {
		case <-parentCtx.Done():
			return parentCtx.Err()
		default:
			_, byteArray, err := conn.Read(parentCtx)
			if err != nil {
				return err
			}

			packetArray, err := manager.DecodeInputs(byteArray)
			if err != nil {
				return err
			}

			err = c.DecodeForListeners(packetArray)
			if err != nil {
				fmt.Printf("Error with the decode for listeners %v\n", err)
				return err
			}
		}
	}
}

func (c *Client) Active() bool {
	return c.connected
}

func (c *Client) Send(packetArray []packet.Packet) {
	c.packetsOut <- packetArray
}

func (c *Client) Close() {
	c.connected = false
	if c.cancel != nil {
		var cancel = c.cancel
		c.cancel = nil
		cancel()
	}
}

// EncodeSystems allows the use of a per-client visibility filter
// so that each client receives a stream of bytes unique to itself
func (c *Client) EncodeSystems() []packet.Packet {
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
	return packetArray
}

func (c *Client) DecodeForListeners(ps []packet.Packet) error {
	var manager = c.manager

	for _, p := range ps {
		var channel = p.Channel
		var data = p.Data

		var l, ok = manager.listeners[channel]
		if ok {
			var listener = *l
			decoder, err := manager.BeginDecode(channel, data)
			if err != nil {
				return err
			}
			err = listener.HandleData(c.identifier, decoder)
			if err != nil {
				return err
			}
			manager.EndDecode(channel)
		}
	}
	return nil
}
