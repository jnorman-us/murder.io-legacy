package input

import "github.com/hajimehoshi/ebiten/v2"

type KeyBinds struct {
	moveForward    ebiten.Key
	moveBackward   ebiten.Key
	moveLeft       ebiten.Key
	moveRight      ebiten.Key
	abilitySpecial ebiten.Key
	abilityAttack  ebiten.MouseButton
	abilityRanged  ebiten.MouseButton
}

func LoadSettings() KeyBinds {
	var keyBinds = KeyBinds{
		moveForward:    ebiten.KeyW,
		moveBackward:   ebiten.KeyS,
		moveLeft:       ebiten.KeyA,
		moveRight:      ebiten.KeyD,
		abilitySpecial: ebiten.KeySpace,
		abilityAttack:  ebiten.MouseButtonLeft,
		abilityRanged:  ebiten.MouseButtonRight,
	}
	return keyBinds
}
