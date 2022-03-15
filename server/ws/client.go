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
	sync.Mutex

	identifier types.UserID
	lobby      *Lobby
	codec      *communications.Codec

	cancel        func()
	connected     bool
	receivedFirst bool

	clumps chan communications.Clump
}

func NewClient(id types.UserID, m *Lobby) *Client {
	return &Client{
		identifier: id,
		lobby:      m,
		clumps:     make(chan communications.Clump),

		connected:     false,
		receivedFirst: false,
	}
}

func (c *Client) Setup(ctx context.Context, conn *websocket.Conn) error {
	fmt.Printf("\"%s\" is connecting.\n", c.identifier)
	c.connected = true
	defer c.Close()

	c.setupCodec()

	group, grpCtx := errgroup.WithContext(ctx)

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
}

func (c *Client) destroyCodec() {
	c.codec = nil
}

func (c *Client) Write(background context.Context, conn *websocket.Conn) error {
	var codec = c.codec
	for {
		select {
		case clump := <-c.clumps:
			byteArray, err := codec.EncodeOutputs(clump)
			if err != nil {
				return err
			}
			err = conn.Write(background, websocket.MessageBinary, byteArray)
			if err != nil {
				c.disconnect()
				return err
			}
		case <-background.Done():
			return background.Err()
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (c *Client) Read(background context.Context, conn *websocket.Conn) error {
	var codec = c.codec

	for {
		select {
		case <-background.Done():
			return background.Err()
		default:
			_, byteArray, err := conn.Read(background)
			if !c.receivedFirst {
				c.receiveFirst()
				clump := c.lobby.EncodeCatchupSystems(c)
				c.Send(clump)
			}
			if err != nil {
				c.disconnect()
				return err
			}

			clump, err := codec.DecodeInputs(byteArray)
			if err != nil {
				return err
			}

			err = c.lobby.DecodeForListeners(c, clump)
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
	c.Lock()
	defer c.Unlock()
	c.receivedFirst = true
}

func (c *Client) disconnect() {
	c.Lock()
	defer c.Unlock()
	if c.connected {
		close(c.clumps)
	}
	c.connected = false
}

func (c *Client) Send(clump communications.Clump) {
	c.Lock()
	defer c.Unlock()

	if c.Active() {
		c.clumps <- clump
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
