package input

type Manager struct {
	inputables     map[int]*Inputable
	identifierToID map[string]int
}

func NewManager() *Manager {
	return &Manager{
		inputables:     map[int]*Inputable{},
		identifierToID: map[string]int{},
	}
}
