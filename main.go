package main

import (
	"fmt"
	"syscall/js"
)

var document js.Value

func H1() {
	myH1 := document.Call("createElement", "h1")
	myH1.Set("innerText", "Created From WASM")
	document.Get("body").Call("appendChild", myH1)
}

func CreateBox() {
	div := document.Call("createElement", "div")
	div.Set("id", "box")
	style := div.Get("style")
	style.Set("width", js.ValueOf("100px"))
	style.Set("height", js.ValueOf("100px"))
	style.Set("position", js.ValueOf("absolute"))
	style.Set("background", js.ValueOf("red"))
	document.Get("body").Call("appendChild", div)
}

func Box(this js.Value, inputs []js.Value) interface{} {
	key := inputs[0].String()
	div := document.Call("getElementById", "box")
	style := div.Get("style")
	go func() {
		switch key {
		case "ArrowUp":
			rect := div.Call("getBoundingClientRect")
			top := rect.Get("top").Float() - 100
			style.Set("top", js.ValueOf(fmt.Sprintf("%.2fpx", top)))
		case "ArrowDown":
			rect := div.Call("getBoundingClientRect")
			top := rect.Get("top").Float() + 100
			style.Set("top", js.ValueOf(fmt.Sprintf("%.2fpx", top)))
		case "ArrowLeft":
			rect := div.Call("getBoundingClientRect")
			left := rect.Get("left").Float() - 100
			style.Set("left", js.ValueOf(fmt.Sprintf("%.2fpx", left)))
		case "ArrowRight":
			rect := div.Call("getBoundingClientRect")
			left := rect.Get("left").Float() + 100
			style.Set("left", js.ValueOf(fmt.Sprintf("%.2fpx", left)))
		}
	}()
	return nil
}

func init() {
	document = js.Global().Get("document")
}

func main() {
	ch := make(chan struct{})
	fmt.Println("============Memulai============")
	H1()
	CreateBox()
	js.Global().Set("Box", js.FuncOf(Box))
	<-ch
}
