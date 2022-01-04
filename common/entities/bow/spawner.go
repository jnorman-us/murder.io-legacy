package bow

type Spawner interface {
	RemoveBow(int)
	SpawnArrow(*Holder, float64)
}

func (b *Bow) SetSpawner(s *Spawner) {
	b.spawner = s
}
