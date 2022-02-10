package game

import "github.com/josephnormandev/murder/common/types"

func (g *Game) SetPlayers(players []types.UserID) {
	for _, player := range players {
		g.Players[player] = 0
	}
}

func (g *Game) ContainsPlayer(player types.UserID) bool {
	var _, ok = g.Players[player]
	return ok
}
