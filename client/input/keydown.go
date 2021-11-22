package input

import (
	"fmt"
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

	if keyCode.Equal(i.keySettings.moveForward) {
		fmt.Println("Moving Forwards")
	} else if keyCode.Equal(i.keySettings.moveLeft) {
		fmt.Println("Moving Left")
	} else if keyCode.Equal(i.keySettings.moveRight) {
		fmt.Println("Moving Right")
	} else if keyCode.Equal(i.keySettings.moveBackward) {
		fmt.Println("Moving Backward")
	}
}
