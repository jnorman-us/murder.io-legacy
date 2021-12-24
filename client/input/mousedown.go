package input

import "syscall/js"

func registerMouseDownListener(i *Manager) {
	i.document.Set("onmousedown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouseDownListener(i, args)
		return nil
	}))
}

func mouseDownListener(i *Manager, args []js.Value) {
	var event = args[0]
	var keyCode = event.Get("which")
	i.updatePlayerInput(keyCode.Int(), true)
}
