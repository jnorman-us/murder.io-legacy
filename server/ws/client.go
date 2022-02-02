package ws

import (
	"context"
	"fmt"
	"github.com/josephnormandev/murder/common/communications"
	"golang.org/x/sync/errgroup"
	"nhooyr.io/websocket"
	"time"
)

type Client struct {
	identifier string
	manager    *Manager
	codec      *communications.Codec

	cancel    func()
	connected bool

	packetsOut chan communications.PacketCollection
}

func NewClient(id string, m *Manager) *Client {
	return &Client{
		identifier: id,
		manager:    m,
		packetsOut: make(chan communications.PacketCollection),
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
	c.codec = communications.NewCodec()

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

func (c *Client) Send(pc communications.PacketCollection) {
	c.packetsOut <- pc
}

func (c *Client) Close() {
	fmt.Printf("\"%s\" is disconnecting.\n", c.identifier)
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
func (c *Client) EncodeSystems() (communications.PacketCollection, error) {
	c.manager.systemMutex.Lock()
	c.manager.spawnMutex.Lock()
	defer c.manager.systemMutex.Unlock()
	defer c.manager.spawnMutex.Unlock()

	var packetArray []communications.Packet
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
			return communications.PacketCollection{}, err
		}
		var outputBytes = codec.EndEncode(channel)

		packetArray = append(packetArray, communications.Packet{
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
			return communications.PacketCollection{}, err
		}
		var outputBytes = codec.EndEncode(class)

		var packet = communications.Packet{
			ID:      id,
			Channel: class,
			Data:    outputBytes,
		}

		packetArray = append(packetArray, packet)
	}
	return communications.PacketCollection{
		PacketArray: packetArray,
		Timestamp:   manager.timestamp,
	}, nil
}

func (c *Client) DecodeForListeners(pc communications.PacketCollection) error {
	var manager = c.manager
	var codec = c.codec

	for _, p := range pc.PacketArray {
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
