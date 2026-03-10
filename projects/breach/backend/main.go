package main

import "syscall/js"

func main() {
	js.Global().Set("runCommand", js.FuncOf(runCommand))
	select {}
}

func runCommand(this js.Value, args []js.Value) any {
	cmd := args[0].String()
	callback := args[1]

	go func() {
		handleCommand(cmd, callback)
	}()

	return nil
}

func handleCommand(cmd string, callback js.Value) {
	callback.Invoke(cmd, "green")
}
