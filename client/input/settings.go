package input

type KeyBinds struct {
	moveForward    int
	moveBackward   int
	moveLeft       int
	moveRight      int
	abilitySpecial int
	abilityAttack  int
	abilityRanged  int
}

func LoadSettings() KeyBinds {
	var keyBinds = KeyBinds{
		moveForward:    87,
		moveBackward:   83,
		moveLeft:       65,
		moveRight:      68,
		abilitySpecial: 32,
		abilityAttack:  0,
		abilityRanged:  1,
	}
	return keyBinds
}
