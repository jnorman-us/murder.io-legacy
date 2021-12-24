package innocent

import (
	"fmt"
	"github.com/josephnormandev/murder/client/input"
)

func (i *Innocent) HandleInputStateChange(s input.Input) {
	fmt.Println(s)
}

func (i *Innocent) AddInputs(m *input.Manager) {
	var listener = input.Listener(i)
	m.AddPlayerListener(&listener)
}
