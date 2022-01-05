package arrow

type Spawner interface {
	RemoveArrow(int)
	RemoveArrowCollidable(int)
}

func (a *Arrow) SetSpawner(s *Spawner) {
	a.spawner = s
}
