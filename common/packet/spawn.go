package packet

import (
	"bytes"
	"encoding/gob"
)

// Spawn is the interface that entities use to encode
// their own data
type Spawn interface {
	GetID() int
	GetClass() string
	GetData(*gob.Encoder)
}

func (m *Manager) AddSpawn(id int, s *Spawn) {
	m.Spawns[id] = s

	var class = (*s).GetClass()
	if _, ok := m.outputs[class]; !ok {
		var classOutput = new(bytes.Buffer)
		m.outputs[class] = classOutput
		m.encoders[class] = gob.NewEncoder(classOutput)
	}
}

func (m *Manager) RemoveSpawn(id int) {
	delete(m.Spawns, id)
}
