package input

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/types"
)

type Manager struct {
	inputs  types.Input
	current map[int]bool
}

func NewManager() *Manager {
	var input = &Manager{
		inputs: types.Input{},
	}
	return input
}

func (m *Manager) SetInputs(inputs types.Input) {
	m.inputs = inputs
}

func (m *Manager) GetChannel() byte {
	return 0x02
}

func (m *Manager) GetData(e *gob.Encoder) error {
	err := e.Encode(m.inputs)
	return err
}
