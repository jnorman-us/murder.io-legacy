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
	i.SetColor(types.Colors.Green)
	m.AddPlayerListener(&listener)
}
