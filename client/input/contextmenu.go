package input

import "syscall/js"

func registerContextMenuDisabler(i *Manager) {
	i.document.Set("oncontextmenu", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return false
	}))
}
