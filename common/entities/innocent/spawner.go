package innocent

// Spawner is an abstraction of the world
type Spawner interface {
	AddInnocent(*Innocent) int
	RemoveInnocent(int)
	SpawnSword(*Innocent) *Swingable
}

func (i *Innocent) SetSpawner(s Spawner) {
	i.spawner = s
	i.id = s.AddInnocent(i)
}
