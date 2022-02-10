package input

import "github.com/josephnormandev/murder/common/types"

type Sizeable interface {
	GetDimensions() types.Vector
}

func (m *Manager) SetSizeable(s *Sizeable) {
	m.sizeable = s
}
