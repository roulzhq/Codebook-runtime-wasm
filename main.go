package main

import (
	"syscall/js"

	"github.com/robertkrimen/otto"
)

var document = js.Global().Get("document")
var vm = otto.New()

func getElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

func CreateVm() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		vm := map[string]interface{}{
			"execute": js.FuncOf(func(this js.Value, p []js.Value) interface{} {
				input := p[0].String()
				value, err := vm.Run(input)

				if err != nil {
					return js.ValueOf(err.Error())
				}

				s, _ := value.ToString()

				return js.ValueOf(s)
			}),
			"clear": js.FuncOf(func(this js.Value, p []js.Value) interface{} {
				vm = otto.New()

				return true
			}),
		}

		return vm
	})
}

func main() {
	quit := make(chan struct{}, 0)

	js.Global().Set("wasm_create_vm", CreateVm())

	<-quit
}
