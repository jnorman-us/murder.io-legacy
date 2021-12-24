package input

import "syscall/js"

func registerMouseUpListener(i *Manager) {
	i.document.Set("onmouseup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouseUpListener(i, args)
		return nil
	}))
}

func mouseUpListener(i *Manager, args []js.Value) {
	var event = args[0]
	var keyCode = event.Get("which")
	i.updatePlayerInput(keyCode.Int(), false)
}
