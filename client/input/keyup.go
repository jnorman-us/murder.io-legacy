package input

import (
	"syscall/js"
)

func registerKeyUpListener(i *Manager) {
	i.document.Set("onkeyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		keyUpListener(i, args)
		return nil
	}))
}

func keyUpListener(i *Manager, args []js.Value) {
	var event = args[0]
	var keyCode = event.Get("keyCode")
	i.updatePlayerInput(keyCode.Int(), false)
}
