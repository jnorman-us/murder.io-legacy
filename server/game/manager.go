package game

import (
	"github.com/josephnormandev/murder/common/game"
	"github.com/josephnormandev/murder/common/types"
)

type Manager struct {
	games map[types.ID]*game.Game
}
