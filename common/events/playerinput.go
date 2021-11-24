package events

type PlayerInputEvent struct {
	PlayerID string
	Forward  bool
	Backward bool
	Left     bool
	Right    bool
}

func (e *PlayerInputEvent) Dead() bool {
	return !(e.Forward || e.Left || e.Right || e.Backward)
}

type PlayerInputListener interface {
	HandlePlayerInput(PlayerInputEvent)
}

func (m *Manager) RegisterPlayerInputListener(l *PlayerInputListener) {
	m.playerInputListeners = append(m.playerInputListeners, l)
}

func (m *Manager) FirePlayerInputEvent(e PlayerInputEvent) {
	var handler = playerInputHandler{
		event:     e,
		listeners: m.playerInputListeners,
	}
	m.eventQueue <- handler
}

type playerInputHandler struct {
	event     PlayerInputEvent
	listeners []*PlayerInputListener
}

func (p playerInputHandler) handle() {
	for _, listener := range p.listeners {
		(*listener).HandlePlayerInput(p.event)
	}
}
