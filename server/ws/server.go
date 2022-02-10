package ws

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/josephnormandev/murder/common/types"
	"net/http"
	"nhooyr.io/websocket"
)

type Server struct {
	lobbies map[types.ID]*Lobby // per world Manager
}

func NewServer() *Server {
	var server = &Server{
		lobbies: map[types.ID]*Lobby{},
	}
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var id = types.UserID(mux.Vars(r)["id"])

	connection, err := websocket.Accept(w, r, nil)
	if err != nil {
		fmt.Printf("Error accepting WS connection! %v\n", err)
	}
	defer (func() {
		_ = connection.Close(websocket.StatusInternalError, "")
		// do nothing with error, probably already closed
	})()

	// find the lobby that the client is in, if any
	for _, l := range s.lobbies {
		var lobby = *l
		var lobbyInfo = *lobby.info
		if lobbyInfo.ContainsPlayer(types.UserID(id)) {
			var client = NewClient(id, l)
			_, ok := lobby.clients[id]

			if !ok {
				lobby.clients[id] = client
				err := client.Setup(r.Context(), connection)
				if err != nil {
					delete(lobby.clients, id)
				}
			} else {
				_ = connection.Close(websocket.StatusPolicyViolation, "Already connected from another location")
			}
			break
		}
	}
	// perhaps here, we'd connect them to the queueing lobby...
	_ = connection.Close(websocket.StatusPolicyViolation, "Username not registered!")
}

func (s *Server) AddLobby(l *Lobby) {
	var id = (*l.info).GetID()
	s.lobbies[id] = l
}
