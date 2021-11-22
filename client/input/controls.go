package input

import "syscall/js"

type keybinds struct {
	moveForward  js.Value
	moveBackward js.Value
	moveLeft     js.Value
	moveRight    js.Value
}

func LoadSettings() keybinds {
	var keySettings = keybinds{
		moveForward:  js.ValueOf(87),
		moveBackward: js.ValueOf(83),
		moveLeft:     js.ValueOf(65),
		moveRight:    js.ValueOf(68),
	}
	return keySettings
}
