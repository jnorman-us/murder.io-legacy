package innocent

import (
	"github.com/josephnormandev/murder/client/input"
	"github.com/josephnormandev/murder/common/types"
)

func (i *Innocent) HandleInputStateChange(s types.Input) {
	i.input = s
}

func (i *Innocent) AddInputs(m *input.Manager) {
	var listener = input.Listener(i)
	m.AddPlayerListener(&listener)
}
