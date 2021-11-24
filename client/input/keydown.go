package input

import (
	"syscall/js"
)

func registerKeyDownListener(i *Input) {
	i.document.Set("onkeydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		keyDownListener(i, args)
		return nil
	}))
}

func keyDownListener(i *Input, args []js.Value) {
	var event = args[0]
	var keyCode = event.Get("keyCode")
	i.updatePlayerInput(Keys(keyCode), true)
}
