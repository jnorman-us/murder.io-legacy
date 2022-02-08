package ws

import "github.com/josephnormandev/murder/common/types"

type LobbyInfo interface {
	GetID() types.ID
	ContainsPlayer(types.UserID) bool
}
