package game

import "github.com/josephnormandev/murder/common/types"

func (g *Game) SetPlayers(ps []types.UserID) {
	for _, userID := range ps {
		g.players[userID] = 0
	}
}

func (g *Game) ContainsPlayer(id types.UserID) bool {
	var _, ok = g.players[id]
	return ok
}
