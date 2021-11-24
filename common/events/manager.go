package events

import "time"

type Manager struct {
	running              bool
	eventQueue           chan eventHandler
	playerInputListeners []*PlayerInputListener
}

func NewManager() *Manager {
	return &Manager{
		eventQueue:           make(chan eventHandler, 100000),
		playerInputListeners: []*PlayerInputListener{},
	}
}

func (m *Manager) Start() {
	m.running = true
	m.run()
}

func (m *Manager) run() {
	for {
		select {
		case handler := <-m.eventQueue:
			handler.handle()
			break
		default:
			time.Sleep(5 * time.Millisecond)
		}
	}
}
