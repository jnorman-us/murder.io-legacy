package ws

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"nhooyr.io/websocket"
	"strings"
	"syscall/js"
	"time"
)

type Client struct {
	manager *Manager
	url     string

	connected bool
	cancel    func()
}

// NewClient accepts a pointer to Manager, and a string with the
// identifier of the current client
func NewClient(m *Manager, id string) *Client {
	var location = js.Global().Get("location")
	var hostname = location.Get("hostname").String()
	var href = location.Get("href").String()
	var port = strings.Split(strings.Split(href, ":")[2], "/")[0]
	var url = fmt.Sprintf("ws://%s:%s/ws/%s", hostname, port, id)
	return &Client{
		manager: m,
		url:     url,
	}
}

func (c *Client) Connect() error {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel
	c.connected = true

	connection, _, err := websocket.Dial(ctx, c.url, nil)
	if err != nil {
		return err
	}
	defer (func() {
		c.Close()
		_ = connection.Close(websocket.StatusNormalClosure, "")
	})()

	group, grpCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return c.Write(grpCtx, connection)
	})
	group.Go(func() error {
		return c.Read(grpCtx, connection)
	})

	err = group.Wait()
	return err
}

func (c *Client) Write(parentCtx context.Context, conn *websocket.Conn) error {
	var manager = c.manager
	for range time.Tick(50 * time.Millisecond) {
		select {
		case <-parentCtx.Done():
			return parentCtx.Err()
		default:
			packetArray, err := manager.EncodeSystems()
			if err != nil {
				return err
			}

			byteArray, err := manager.EncodeOutputs(packetArray)
			if err != nil {
				return err
			}

			err = conn.Write(parentCtx, websocket.MessageBinary, byteArray)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

			packetCollection, err := manager.DecodeInputs(byteArray)
			if err != nil {
				return err
			}

			err = manager.DecodeForListeners(packetCollection)
			if err != nil {
				return err
			}
		}
	}
}

func (c *Client) Active() bool {
	return c.connected
}

func (c *Client) Close() {
	c.connected = false
	if c.cancel != nil {
		var cancel = c.cancel
		c.cancel = nil
		cancel()
	}
}

func (c *Client) Send() {

}

/*
	c, _, err := ws.Dial(ctx, "ws://localhost:8080", nil)
	if err != nil {
		// ...
	}
	defer c.Close(ws.StatusInternalError, "the sky is falling")

	err = wsjson.Write(ctx, c, "hi")
	if err != nil {
		// ...
	}

	c.Close(ws.StatusNormalClosure, "")
*/
