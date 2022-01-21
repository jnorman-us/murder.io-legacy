package ws

import (
	"context"
	"github.com/josephnormandev/murder/common/packet"
	"golang.org/x/sync/errgroup"
	"nhooyr.io/websocket"
	"time"
)

type Client struct {
	identifier string
	manager    *Manager
	codec      *packet.Codec

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

	c.setupCodec()

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

func (c *Client) setupCodec() {
	c.codec = packet.NewCodec()

	for class := range c.manager.classes {
		c.codec.AddEncoder(class)
	}
	for channel := range c.manager.systems {
		c.codec.AddEncoder(channel)
	}
	for channel := range c.manager.listeners {
		c.codec.AddDecoder(channel)
	}
	c.manager.codecs[c.identifier] = c.codec
}

func (c *Client) destroyCodec() {
	c.codec = nil
	delete(c.manager.codecs, c.identifier)
}

func (c *Client) Write(parentCtx context.Context, conn *websocket.Conn) error {
	var codec = c.codec
	for {
		select {
		case packets := <-c.packetsOut:
			byteArray, err := codec.EncodeOutputs(packets)
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
	var codec = c.codec

	for {
		select {
		case <-parentCtx.Done():
			return parentCtx.Err()
		default:
			_, byteArray, err := conn.Read(parentCtx)
			if err != nil {
				return err
			}

			packetArray, err := codec.DecodeInputs(byteArray)
			if err != nil {
				return err
			}

			err = c.DecodeForListeners(packetArray)
			if err != nil {
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
	c.destroyCodec()
	if c.cancel != nil {
		var cancel = c.cancel
		c.cancel = nil
		cancel()
	}
}

// EncodeSystems allows the use of a per-client visibility filter
// so that each client receives a stream of bytes unique to itself
func (c *Client) EncodeSystems() ([]packet.Packet, error) {
	var packetArray []packet.Packet
	var manager = c.manager
	var codec = c.codec

	// perhaps implement this in the manager because
	// there is no need to make this data unique per client
	for _, s := range manager.systems {
		var system = *s
		var channel = system.GetChannel()

		var encoder = codec.BeginEncode(channel)
		err := system.GetData(encoder)
		if err != nil {
			return []packet.Packet{}, err
		}
		var outputBytes = codec.EndEncode(channel)

		packetArray = append(packetArray, packet.Packet{
			ID:      0,
			Channel: channel,
			Data:    outputBytes,
		})
	}

	for _, s := range manager.spawns {
		var spawn = *s
		var id = spawn.GetID()
		var class = spawn.GetClass()

		var encoder = codec.BeginEncode(class)
		err := spawn.GetData(encoder)
		if err != nil {
			return []packet.Packet{}, err
		}
		var outputBytes = codec.EndEncode(class)

		var packet = packet.Packet{
			ID:      id,
			Channel: class,
			Data:    outputBytes,
		}

		packetArray = append(packetArray, packet)
	}
	return packetArray, nil
}

func (c *Client) DecodeForListeners(ps []packet.Packet) error {
	var manager = c.manager
	var codec = c.codec

	for _, p := range ps {
		var channel = p.Channel
		var data = p.Data

		var l, ok = manager.listeners[channel]
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
	}
	return nil
}
