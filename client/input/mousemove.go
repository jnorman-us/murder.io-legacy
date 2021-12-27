package input

import (
	"syscall/js"
)

func registerMouseMoveListener(i *Manager) {
	i.document.Set("onmousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mouseMoveListener(i, args)
		return nil
	}))
}

func mouseMoveListener(i *Manager, args []js.Value) {
	var event = args[0]
	var x = event.Get("clientX").Float()
	var y = event.Get("clientY").Float()
	i.updatePlayerDirection(x, y)
}
