package input

import "syscall/js"

type Input struct {
	window   js.Value
	document js.Value

	keySettings keybinds
}

func NewInput() *Input {
	var input = &Input{}
	input.window = js.Global()
	input.document = input.window.Get("document")
	input.keySettings = LoadSettings()

	registerKeyDownListener(input)
	registerKeyUpListener(input)

	return input
}
