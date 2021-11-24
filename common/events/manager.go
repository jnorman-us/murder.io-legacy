package events

var Singleton *Manager

type Manager struct {
	eventQueue           chan eventHandler
	playerInputListeners []*PlayerInputListener
}

func InitializeEvents() {
	Singleton = &Manager{
		eventQueue:           make(chan eventHandler, 100000),
		playerInputListeners: []*PlayerInputListener{},
	}
}
