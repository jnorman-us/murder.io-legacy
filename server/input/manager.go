package input

import "github.com/josephnormandev/murder/common/types"

type Manager struct {
	inputables     map[int]*Inputable
	identifierToID map[types.UserID]int
}

func NewManager() *Manager {
	return &Manager{
		inputables:     map[int]*Inputable{},
		identifierToID: map[types.UserID]int{},
	}
}
