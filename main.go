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

func executeJs(this js.Value, p []js.Value) interface{} {
	input := p[0].String()

	value, err := vm.Run(input)

	if err != nil {
		return js.ValueOf(err.Error())
	}

	s, _ := value.ToString()

	return js.ValueOf(s)
}

func main() {
	quit := make(chan struct{}, 0)

	js.Global().Set("cb_execute", js.FuncOf(executeJs))

	<-quit
}
