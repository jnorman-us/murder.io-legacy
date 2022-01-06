package collisions

type Manager struct {
	ArrowWalls   map[int]*ArrowWall
	WallArrows   map[int]*WallArrow
	ArrowPlayers map[int]*ArrowPlayer
	PlayerArrows map[int]*PlayerArrow
	PlayerWalls  map[int]*PlayerWall
	WallPlayers  map[int]*WallPlayer
	SwordPlayers map[int]*SwordPlayer
	PlayerSwords map[int]*PlayerSword
}

func NewManager() *Manager {
	return &Manager{
		ArrowWalls:   map[int]*ArrowWall{},
		WallArrows:   map[int]*WallArrow{},
		ArrowPlayers: map[int]*ArrowPlayer{},
		PlayerArrows: map[int]*PlayerArrow{},
		PlayerWalls:  map[int]*PlayerWall{},
		WallPlayers:  map[int]*WallPlayer{},
		SwordPlayers: map[int]*SwordPlayer{},
		PlayerSwords: map[int]*PlayerSword{},
	}
}

func (m *Manager) ResolveCollisions() {
	m.resolvePlayerWalls()
	m.resolveSwordPlayers()
	m.resolveArrowPlayers()
	m.resolveArrowWalls()
}
