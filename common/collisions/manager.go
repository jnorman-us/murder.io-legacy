package collisions

type Manager struct {
	PlayerPlayerCollidables map[int]*PlayerPlayerCollidable
	ArrowWalls              map[int]*ArrowWall
	WallArrows              map[int]*WallArrow
	ArrowPlayers            map[int]*ArrowPlayer
	PlayerArrows            map[int]*PlayerArrow
}

func NewManager() *Manager {
	return &Manager{
		PlayerPlayerCollidables: map[int]*PlayerPlayerCollidable{},
		ArrowWalls:              map[int]*ArrowWall{},
		WallArrows:              map[int]*WallArrow{},
		ArrowPlayers:            map[int]*ArrowPlayer{},
		PlayerArrows:            map[int]*PlayerArrow{},
	}
}

func (m *Manager) ResolveCollisions() {
	m.resolvePlayerPlayerCollidables()
	m.resolveArrowWalls()
	m.resolveArrowPlayers()
}
