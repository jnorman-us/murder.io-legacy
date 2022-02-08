package input

import "github.com/josephnormandev/murder/common/types"

type Inputable interface {
	GetID() int
	GetIdentifier() types.UserID
	SetInput(types.Input)
}

func (m *Manager) AddPlayerListener(id int, i *Inputable) {
	var identifier = (*i).GetIdentifier()

	m.inputables[id] = i
	m.identifierToID[identifier] = id
}

func (m *Manager) RemovePlayerListener(i int) {
	delete(m.inputables, i)

	for identifier, id := range m.identifierToID {
		if id == i {
			delete(m.identifierToID, identifier)
		}
	}
}
