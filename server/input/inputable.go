package input

import "github.com/josephnormandev/murder/common/types"

type Inputable interface {
	GetID() types.ID
	GetUserID() types.UserID
	SetInput(types.Input)
}

func (m *Manager) AddPlayerListener(id types.ID, i *Inputable) {
	var identifier = (*i).GetUserID()

	m.inputables[id] = i
	m.identifierToID[identifier] = id
}

func (m *Manager) RemovePlayerListener(i types.ID) {
	delete(m.inputables, i)

	for identifier, id := range m.identifierToID {
		if id == i {
			delete(m.identifierToID, identifier)
		}
	}
}
