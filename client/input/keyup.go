package input

import (
	"fmt"
	"syscall/js"
)

func registerKeyUpListener(i *Input) {
	i.document.Set("onkeyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		keyUpListener(i, args)
		return nil
	}))
}

func keyUpListener(i *Input, args []js.Value) {
	var event = args[0]
	var keyCode = event.Get("keyCode")

	if keyCode.Equal(i.keySettings.moveForward) {
		fmt.Println("Done Moving Forwards")
	} else if keyCode.Equal(i.keySettings.moveLeft) {
		fmt.Println("Done Moving Left")
	} else if keyCode.Equal(i.keySettings.moveRight) {
		fmt.Println("Done Moving Right")
	} else if keyCode.Equal(i.keySettings.moveBackward) {
		fmt.Println("Done Moving Backward")
	}
}
