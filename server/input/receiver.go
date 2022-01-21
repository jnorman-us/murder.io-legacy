package input

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/types"
)

func (m *Manager) GetChannel() string {
	return "INPUT"
}

func (m *Manager) HandleData(identifier string, decoder *gob.Decoder) error {
	var input = &types.Input{}
	err := decoder.Decode(input)
	if err != nil {
		return err
	}

	var id, ok = m.identifierToID[identifier]
	if ok {
		var inputable = *m.inputables[id]
		inputable.HandleInputStateChange(*input)
	}
	return nil
}
