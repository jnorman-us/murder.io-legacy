package ws

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"nhooyr.io/websocket"
	"time"
)

type Server struct {
	manager *Manager
	clients map[string]*Client
}

func NewServer(identifiers []string, manager *Manager) *Server {
	var clients = map[string]*Client{}
	for _, identifier := range identifiers {
		clients[identifier] = NewClient(identifier, manager)
	}

	var server = &Server{
		manager: manager,
		clients: clients,
	}
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var id = mux.Vars(r)["id"]

	fmt.Printf("ID: %s\n", id)
	connection, err := websocket.Accept(w, r, nil)
	if err != nil {
		fmt.Printf("Error accepting WS connection! %v\n", err)
	}
	defer (func() {
		err := connection.Close(websocket.StatusInternalError, "")
		if err != nil {
			fmt.Printf("Error closing WS connection! %v\n", err)
		}
	})()

	var client = s.clients[id]
	err = client.Setup(r.Context(), connection)
	if err != nil {
		fmt.Printf("Error with WS! %v\n", err)
	}
}

func (s *Server) Send() {
	for range time.Tick(50 * time.Millisecond) {
		for _, c := range s.clients {
			var client = *c
			if client.Active() {
				var packetArray = client.EncodeSystems()
				client.Send(packetArray)
			}
		}
	}
}
