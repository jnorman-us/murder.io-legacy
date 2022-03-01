package input

import (
	"encoding/gob"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josephnormandev/murder/common/types"
)

type Manager struct {
	keyBinds KeyBinds

	inputs  types.Input
	current map[int]bool
}

func NewManager() *Manager {
	var input = &Manager{
		keyBinds: LoadSettings(),

		inputs:  types.Input{},
		current: map[int]bool{},
	}
	return input
}

func (m *Manager) PollInputs() {
	// var mouseX, mouseY = ebiten.CursorPosition()

	var keyBinds = m.keyBinds
	var inputs = types.Input{}

	inputs.Left = ebiten.IsKeyPressed(keyBinds.moveLeft)
	inputs.Right = ebiten.IsKeyPressed(keyBinds.moveRight)
	inputs.Forward = ebiten.IsKeyPressed(keyBinds.moveForward)
	inputs.Backward = ebiten.IsKeyPressed(keyBinds.moveBackward)
	inputs.Special = ebiten.IsKeyPressed(keyBinds.abilitySpecial)
	inputs.AttackClick = ebiten.IsMouseButtonPressed(keyBinds.abilityAttack)
	inputs.RangedClick = ebiten.IsMouseButtonPressed(keyBinds.abilityRanged)

	if !inputs.Equals(m.inputs) {
		m.inputs = inputs
	}
}

func (m *Manager) GetChannel() byte {
	return 0x02
}

func (m *Manager) GetData(e *gob.Encoder) error {
	err := e.Encode(m.inputs)
	return err
}
