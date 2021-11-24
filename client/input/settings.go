package input

import "syscall/js"

type Keys js.Value

type keybinds struct {
	moveForward  Keys
	moveBackward Keys
	moveLeft     Keys
	moveRight    Keys
}

func LoadSettings() keybinds {
	var keySettings = keybinds{
		moveForward:  Keys(js.ValueOf(87)),
		moveBackward: Keys(js.ValueOf(83)),
		moveLeft:     Keys(js.ValueOf(65)),
		moveRight:    Keys(js.ValueOf(68)),
	}
	return keySettings
}

func (k Keys) equals(o Keys) bool {
	return js.Value(k).Equal(js.Value(o))
}
