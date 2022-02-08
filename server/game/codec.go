package game

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/server/ws"
)

func (g *ServerGame) GetChannel() string {
	return "s"
}

func (g *ServerGame) Flush() {

}

func (g *ServerGame) GetData(encoder *gob.Encoder) error {
	return encoder.Encode(g)
}

func (g *ServerGame) GetLobby() *ws.Lobby {
	return g.packets
}
