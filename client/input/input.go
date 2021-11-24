package input

import (
	"github.com/josephnormandev/murder/common/events"
	"syscall/js"
)

type Input struct {
	eventsManager *events.Manager
	window        js.Value
	document      js.Value

	keySettings keybinds

	playerInput events.PlayerInputEvent
}

func NewInput(e *events.Manager, playerID string) *Input {
	var input = &Input{}
	input.eventsManager = e
	input.window = js.Global()
	input.document = input.window.Get("document")
	input.keySettings = LoadSettings()
	input.playerInput = events.PlayerInputEvent{}

	registerKeyDownListener(input)
	registerKeyUpListener(input)

	return input
}

func (i *Input) updatePlayerInput(key Keys, active bool) {
	var newInput = i.playerInput
	if key.equals(i.keySettings.moveLeft) {
		if i.playerInput.Right == true && active {
			return
		}
		newInput.Left = active
	}
	if key.equals(i.keySettings.moveRight) {
		if i.playerInput.Left == true && active {
			return
		}
		newInput.Right = active
	}
	if key.equals(i.keySettings.moveForward) {
		if i.playerInput.Backward == true && active {
			return
		}
		newInput.Forward = active
	}
	if key.equals(i.keySettings.moveBackward) {
		if i.playerInput.Forward == true && active {
			return
		}
		newInput.Backward = active
	}
	if newInput != i.playerInput {
		i.playerInput = newInput
		i.eventsManager.FirePlayerInputEvent(i.playerInput)
	}
}
