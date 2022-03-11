package ws

import (
	"context"
	"fmt"
	"github.com/josephnormandev/murder/common/communications"
	"github.com/josephnormandev/murder/common/types"
	"golang.org/x/sync/errgroup"
	"nhooyr.io/websocket"
	"sync"
	"time"
)

type Client struct {
	identifier types.UserID
	lobby      *Lobby
	codec      *communications.Codec

	cancel        func()
	connected     bool
	receivedFirst bool

	packetsOut  chan communications.PacketCollection
	channelLock sync.Mutex
}

func NewClient(id types.UserID, m *Lobby) *Client {
	return &Client{
		identifier: id,
		lobby:      m,
		packetsOut: make(chan communications.PacketCollection),

		connected:     false,
		receivedFirst: false,
	}
}

func (c *Client) Setup(ctx context.Context, conn *websocket.Conn) error {
	fmt.Printf("\"%s\" is connecting.\n", c.identifier)
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

	for class := range c.lobby.classes {
		c.codec.AddEncoder(class)
	}
	for channel := range c.lobby.systems {
		c.codec.AddEncoder(channel)
	}
	for channel := range c.lobby.listeners {
		c.codec.AddDecoder(channel)
	}
}

func (c *Client) destroyCodec() {
	c.codec = nil
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
				c.disconnect()
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
			if !c.receivedFirst {
				c.receiveFirst()

				packetCollection, err := (*c).EncodeCatchupSystems()
				if err != nil {
					return err
				}
				c.Send(packetCollection)
			}
			if err != nil {
				c.disconnect()
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
	return c.connected && c.receivedFirst
}

func (c *Client) receiveFirst() {
	c.channelLock.Lock()
	defer c.channelLock.Unlock()
	c.receivedFirst = true
}

func (c *Client) disconnect() {
	c.channelLock.Lock()
	defer c.channelLock.Unlock()
	if c.connected {
		close(c.packetsOut)
	}
	c.connected = false
}

func (c *Client) Send(pc communications.PacketCollection) {
	c.channelLock.Lock()
	defer c.channelLock.Unlock()
	if c.Active() {
		c.packetsOut <- pc
	}
}

func (c *Client) Close() {
	fmt.Printf("\"%s\" is disconnecting.\n", c.identifier)
	c.disconnect()
	c.destroyCodec()
	if c.cancel != nil {
		c.cancel()
	}
}

// EncodeSystems allows the use of a per-client visibility filter
// so that each client receives a stream of bytes unique to itself
func (c *Client) EncodeSystems() (communications.PacketCollection, error) {
	c.lobby.systemMutex.Lock()
	c.lobby.spawnMutex.Lock()
	defer c.lobby.systemMutex.Unlock()
	defer c.lobby.spawnMutex.Unlock()

	var packetArray []communications.Packet
	var lobby = c.lobby
	var codec = c.codec

	// perhaps implement this in the manager because
	// there is no need to make this data unique per client
	for _, s := range lobby.systems {
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

	for _, s := range lobby.spawns {
		var spawn = *s
		if !spawn.Dirty() {
			continue
		}
		spawn.CleanDirt()

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
		Timestamp:   lobby.time.Tick,
	}, nil
}

func (c *Client) EncodeCatchupSystems() (communications.PacketCollection, error) {
	c.lobby.systemMutex.Lock()
	c.lobby.spawnMutex.Lock()
	defer c.lobby.systemMutex.Unlock()
	defer c.lobby.spawnMutex.Unlock()

	var packetArray []communications.Packet
	var lobby = c.lobby
	var codec = c.codec

	// perhaps implement this in the manager because
	// there is no need to make this data unique per client
	for _, s := range lobby.systems {
		var system = *s
		var channel = system.GetChannel()

		var encoder = codec.BeginEncode(channel)
		err := system.GetCatchupData(encoder)
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

	for _, s := range lobby.spawns {
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
		Timestamp:   lobby.time.Tick,
	}, nil
}

func (c *Client) DecodeForListeners(pc communications.PacketCollection) error {
	var lobby = c.lobby
	var codec = c.codec

	for _, p := range pc.PacketArray {
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
	}
	return nil
}
