package tcp

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Server struct {
	packetManager *Manager

	server   *net.TCPListener
	clients  map[string]*Client
	uSockets map[int]*net.TCPConn
	socketI  int
}

func NewServer(manager *Manager, identifiers []string) *Server {
	var addr, _ = net.ResolveTCPAddr("tcp", ":30300")
	var server, err = net.ListenTCP("tcp", addr)

	if err != nil {
		fmt.Printf("Error starting server! %v\n", err)
		return nil
	}

	var clients = map[string]*Client{}
	for _, identifier := range identifiers {
		clients[identifier] = NewClient(identifier, manager)
	}

	return &Server{
		packetManager: manager,

		server:   server,
		clients:  clients,
		uSockets: map[int]*net.TCPConn{},
	}
}

func (s *Server) AcceptConnections() {
	var server = s.server

	for {
		var socket, err = server.AcceptTCP()

		fmt.Printf("Test")

		if err != nil {
			fmt.Printf("Error accepting connection! %v\n", err)
		} else {
			s.uSockets[s.socketI] = socket
			s.socketI++
		}
	}
}

func (s *Server) Listen() {
	var packetManager = s.packetManager
	var inputBuffer = packetManager.GetInputBuffer()

	for range time.Tick(5 * time.Millisecond) {
		for _, c := range s.clients {
			var client = *c
			var reader = client.reader

			if client.Connected() {
				inputBuffer.Reset()
				var n, err = reader.WriteTo(inputBuffer)

				if n == 0 {
					continue
				}

				if err != nil {
					fmt.Printf("Error reading packet! %v\n", err)
				} else {
					var _, packets = packetManager.DecodeInputs()
					client.DecodeForListeners(packets)
				}
			}
		}

		for i, socket := range s.uSockets {
			var reader = bufio.NewReader(socket)

			inputBuffer.Reset()
			var n, err = reader.WriteTo(inputBuffer)

			if n == 0 {
				continue
			}

			if err != nil {
				fmt.Printf("Error reading packet from UFO! %v\n", err)
			} else {
				var ident, packets = packetManager.DecodeInputs()
				var client, ok = s.clients[ident]

				if ok {
					client.SetSocket(socket)
					client.DecodeForListeners(packets)
					delete(s.uSockets, i)
				} else {
					socket.Close()
					delete(s.uSockets, i)
				}
			}
		}
	}
}

func (s *Server) Send() {
	var manager = s.packetManager
	var outputBuffer = manager.GetOutputBuffer()
	for range time.Tick(1000 * time.Millisecond) {
		for _, c := range s.clients {
			var client = *c
			var socket = client.GetSocket()

			if client.Connected() {
				outputBuffer.Reset()
				client.EncodeSystems()

				var _, err = socket.ReadFrom(outputBuffer)

				if err != nil {
					fmt.Printf("Error writing packet! %v\n", err)
				}
			}
		}
	}
}

func (s *Server) Close() {
	var server = *s.server
	var err = server.Close()

	if err != nil {
		fmt.Printf("Error closing server! %v\n", err)
	}
}

/*

func main() {
    p := make([]byte, 2048)
    addr := net.UDPAddr{
        Port: 1234,
        IP: net.ParseIP("127.0.0.1"),
    }
    ser, err := net.ListenUDP("tcp", &addr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }
    for {
        _,remoteaddr,err := ser.ReadFromUDP(p)
        fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
        if err !=  nil {
            fmt.Printf("Some error  %v", err)
            continue
        }
        go sendResponse(ser, remoteaddr)
    }
}

*/
