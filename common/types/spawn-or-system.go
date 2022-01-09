package types

type SpawnOrSystem bool

func (s SpawnOrSystem) IsSpawn() bool {
	return s == true
}

func (s SpawnOrSystem) IsSystem() bool {
	return s == false
}

func Spawn() SpawnOrSystem {
	return true
}

func System() SpawnOrSystem {
	return false
}
