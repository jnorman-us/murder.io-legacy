package input

type Listener interface {
	HandleInputStateChange(Input)
}

func (m *Manager) AddPlayerListener(l *Listener) {
	m.playerListener = l
}

func (m *Manager) RemovePlayerListener() {
	m.playerListener = nil
}
