package innocent

// Spawner is an abstraction of the world
type Spawner interface {
	RemoveInnocent(int)
	SpawnSword(*Innocent) *Swingable
}

func (i *Innocent) SetSpawner(s *Spawner) {
	i.spawner = s
}
