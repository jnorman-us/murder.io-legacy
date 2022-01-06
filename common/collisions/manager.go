package collisions

type Manager struct {
	ArrowWalls   map[int]*ArrowWall
	WallArrows   map[int]*WallArrow
	ArrowPlayers map[int]*ArrowPlayer
	PlayerArrows map[int]*PlayerArrow
	PlayerWalls  map[int]*PlayerWall
	WallPlayers  map[int]*WallPlayer
}

func NewManager() *Manager {
	return &Manager{
		ArrowWalls:   map[int]*ArrowWall{},
		WallArrows:   map[int]*WallArrow{},
		ArrowPlayers: map[int]*ArrowPlayer{},
		PlayerArrows: map[int]*PlayerArrow{},
		PlayerWalls:  map[int]*PlayerWall{},
		WallPlayers:  map[int]*WallPlayer{},
	}
}

func (m *Manager) ResolveCollisions() {
	m.resolvePlayerWalls()
	m.resolveArrowPlayers()
	m.resolveArrowWalls()
}
