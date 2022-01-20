package input

import (
	"encoding/gob"
	"fmt"
	"github.com/josephnormandev/murder/common/types"
)

func (m *Manager) GetChannel() string {
	return "INPUT"
}

func (m *Manager) HandleData(identifier string, decoder *gob.Decoder) error {
	var input = &types.Input{}
	err := decoder.Decode(input)
	if err != nil {
		fmt.Printf("ERror with the decode! %v\n", err)
		return err
	}

	fmt.Println(identifier, input)
	return nil
}
