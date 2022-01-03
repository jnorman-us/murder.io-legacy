package input

import "github.com/josephnormandev/murder/common/types"

type Listener interface {
	HandleInputStateChange(types.Input)
}

func (m *Manager) AddPlayerListener(l *Listener) {
	m.playerListener = l
}

func (m *Manager) RemovePlayerListener() {
	m.playerListener = nil
}
