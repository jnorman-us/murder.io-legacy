package input

import (
	"encoding/gob"
	"fmt"
	"github.com/josephnormandev/murder/common/types"
	"syscall/js"
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

func (m *Manager) SetInputs(this js.Value, values []js.Value) interface{} {
	fmt.Println(values)
	return nil
}

func (m *Manager) GetChannel() byte {
	return 0x02
}

func (m *Manager) GetData(e *gob.Encoder) error {
	err := e.Encode(m.inputs)
	return err
}
