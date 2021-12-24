package input

import (
	"syscall/js"
)

type Manager struct {
	window   js.Value
	document js.Value
	canvas   js.Value
	keyBinds KeyBinds

	playerListener *Listener

	inputs  Input
	current map[int]bool
}

func NewManager(playerID string) *Manager {
	var input = &Manager{
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
				newInputs.forward = true
				newInputs.backward = false
			} else {
				newInputs.forward = false
				newInputs.backward = true
			}
		} else {
			newInputs.forward = active
		}
	} else if key == i.keyBinds.moveBackward {
		if _, ok := i.current[i.keyBinds.moveForward]; ok {
			if active {
				newInputs.backward = true
				newInputs.forward = false
			} else {
				newInputs.backward = false
				newInputs.forward = true
			}
		} else {
			newInputs.backward = active
		}
	}
	if key == i.keyBinds.moveLeft {
		if _, ok := i.current[i.keyBinds.moveRight]; ok {
			if active {
				newInputs.left = true
				newInputs.right = false
			} else {
				newInputs.left = false
				newInputs.right = true
			}
		} else {
			newInputs.left = active
		}
	} else if key == i.keyBinds.moveRight {
		if _, ok := i.current[i.keyBinds.moveLeft]; ok {
			if active {
				newInputs.right = true
				newInputs.left = false
			} else {
				newInputs.right = false
				newInputs.left = true
			}
		} else {
			newInputs.right = active
		}
	}
	/*
		if key.equals(i.keyBinds.moveLeft) {
		} else if key.equals(i.keyBinds.moveRight) {
		}
		if key.equals(i.keyBinds.abilityAttack) {

		}
		if key.equals(i.keyBinds.abilityRanged) {

		}
		if key.equals(i.keyBinds.abilitySpecial) {

		}*/
	if !i.inputs.Equals(newInputs) {
		i.inputs = newInputs
		if i.playerListener != nil {
			(*i.playerListener).HandleInputStateChange(newInputs)
		}
	}
}
