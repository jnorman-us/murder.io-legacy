package innocent

import (
	"github.com/josephnormandev/murder/client/input"
)

func (i *Innocent) HandleInputStateChange(s input.Input) {
	i.input = s
}

func (i *Innocent) AddInputs(m *input.Manager) {
	var listener = input.Listener(i)
	m.AddPlayerListener(&listener)
}