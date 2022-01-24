package deletions

import (
	"encoding/gob"
	"sync"
)

type Manager struct {
	removedIDs map[int]int
	flushedIDs map[int]int
	sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		removedIDs: map[int]int{},
		flushedIDs: map[int]int{},
	}
}

func (m *Manager) GetChannel() string {
	return "dlt"
}

func (m *Manager) Flush() {
	m.Lock()
	defer m.Unlock()

	m.flushedIDs = map[int]int{}
	for id := range m.removedIDs {
		m.flushedIDs[id] = 0
	}
	m.removedIDs = map[int]int{}
}

func (m *Manager) RemoveID(id int) {
	m.Lock()
	defer m.Unlock()
	m.removedIDs[id] = 0
}

func (m *Manager) GetData(encoder *gob.Encoder) error {
	return encoder.Encode(m.flushedIDs)
}
