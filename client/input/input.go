package input

import (
	"github.com/josephnormandev/murder/common/types"
	"syscall/js"
)

type Manager struct {
	window   js.Value
	document js.Value
	sizeable *Sizeable

	keyBinds KeyBinds

	playerListener *Listener

	inputs  types.Input
	current map[int]bool
}

func NewManager(s *Sizeable) *Manager {
	var input = &Manager{
		sizeable: s,
		window:   js.Global(),
		current:  map[int]bool{},
		keyBinds: LoadSettings(),
	}
	input.document = input.window.Get("document")

	registerKeyDownListener(input)
	registerKeyUpListener(input)
	registerMouseUpListener(input)
	registerMouseDownListener(input)
	registerContextMenuDisabler(input)
	registerMouseMoveListener(input)

	return input
}

func (i *Manager) updatePlayerInput(key int, active bool) {
	var newInputs = i.inputs

	if _, ok := i.current[key]; !ok && active {
		i.current[key] = true
	} else {
		delete(i.current, key)
	}

	if key == i.keyBinds.moveForward {
		if _, ok := i.current[i.keyBinds.moveBackward]; ok {
			if active {
				newInputs.Forward = true
				newInputs.Backward = false
			} else {
				newInputs.Forward = false
				newInputs.Backward = true
			}
		} else {
			newInputs.Forward = active
		}
	} else if key == i.keyBinds.moveBackward {
		if _, ok := i.current[i.keyBinds.moveForward]; ok {
			if active {
				newInputs.Backward = true
				newInputs.Forward = false
			} else {
				newInputs.Backward = false
				newInputs.Forward = true
			}
		} else {
			newInputs.Backward = active
		}
	}
	if key == i.keyBinds.moveLeft {
		if _, ok := i.current[i.keyBinds.moveRight]; ok {
			if active {
				newInputs.Left = true
				newInputs.Right = false
			} else {
				newInputs.Left = false
				newInputs.Right = true
			}
		} else {
			newInputs.Left = active
		}
	} else if key == i.keyBinds.moveRight {
		if _, ok := i.current[i.keyBinds.moveLeft]; ok {
			if active {
				newInputs.Right = true
				newInputs.Left = false
			} else {
				newInputs.Right = false
				newInputs.Left = true
			}
		} else {
			newInputs.Right = active
		}
	}
	if key == i.keyBinds.abilityAttack {
		newInputs.AttackClick = active
	}
	if key == i.keyBinds.abilityRanged {
		newInputs.RangedClick = active
	}
	if !i.inputs.Equals(newInputs) {
		i.inputs = newInputs
		if i.playerListener != nil {
			(*i.playerListener).HandleInputStateChange(newInputs)
		}
	}
}

func (m *Manager) updatePlayerDirection(x, y float64) {
	var mousePosition = types.NewVector(x, y)
	mousePosition.Scale(-1)
	var center = (*m.sizeable).GetDimensions()
	center.Scale(.5)
	center.Add(mousePosition)
	center.Scale(-1)

	var newAngle = center.Angle()

	if newAngle != m.inputs.Direction {
		m.inputs.Direction = center.Angle()
		(*m.playerListener).HandleInputStateChange(m.inputs)
	}
}
