package drifter

import (
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

type State struct {
	types.Change
	entities.Health
	types.UserID
}
