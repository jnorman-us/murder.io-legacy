package input

import "github.com/josephnormandev/murder/common/types"

type Manager struct {
	inputables     map[types.ID]*Inputable
	identifierToID map[types.UserID]types.ID
}

func NewManager() *Manager {
	return &Manager{
		inputables:     map[types.ID]*Inputable{},
		identifierToID: map[types.UserID]types.ID{},
	}
}
