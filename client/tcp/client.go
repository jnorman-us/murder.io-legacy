package tcp

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Client struct {
	connection *net.Conn
	reader     *bufio.Reader
	writer     *bufio.Writer

	packetManager *Manager
}

func NewClient(manager *Manager) *Client {
	var connection, err = net.Dial("tcp", "127.0.0.1:30300")

	if err != nil {
		fmt.Printf("Error connecting to server! %v\n", err)
		return nil
	}

	var reader = bufio.NewReader(connection)
	var writer = bufio.NewWriter(connection)

	return &Client{
		connection: &connection,
		reader:     reader,
		writer:     writer,

		packetManager: manager,
	}
}

func (c *Client) Listen() {
	var packetManager = c.packetManager
	var inputBuffer = packetManager.GetInputBuffer()

	for {
		inputBuffer.Reset()
		var n, err = c.reader.WriteTo(inputBuffer)

		if n == 0 {
			time.Sleep(5 * time.Millisecond)
			continue
		}

		if err != nil {
			fmt.Printf("Error reading packet! %v\n", err)
		} else {
			var _, packets = packetManager.DecodeInputs()
			packetManager.DecodeForListeners(packets)
		}
	}
}

func (c *Client) Send() {
	var packetManager = c.packetManager
	var outputBuffer = packetManager.GetOutputBuffer()

	for range time.Tick(1000 * time.Millisecond) {
		outputBuffer.Reset()
		packetManager.EncodeSystems()
		var _, err = c.writer.ReadFrom(outputBuffer)

		if err != nil {
			fmt.Printf("Error writing packet! %v\n", err)
		}
	}
}

func (c *Client) Close() {
	var connection = *c.connection
	var err = connection.Close()

	if err != nil {
		fmt.Printf("Error closing connection! %v", err)
	}
}
