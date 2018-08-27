package main

import (
	"fmt"
	"github.com/go_wasm/wasm"
	"strconv"
	"syscall/js"
	"time"
)

func add(i []js.Value) {

	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	sum := wasm.Madd(int1, int2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", sum)

	go ticker(i)
	go WebRequest(i)

	println(int1 + int2)

}

func ticker(t []js.Value) {

	for i := 0; i < 13; i++ {
		js.Global().Get("document").Call("getElementById", t[3].String()).Set("value", wasm.Mtext(i))
		time.Sleep(2 * time.Second)
	}

}

func WebRequest(t []js.Value) {
	url := "https://httpbin.org/get"
	js.Global().Get("document").Call("getElementById", t[4].String()).Set("value",
		wasm.MakeRequest(url))

}

func registerCallbacks() {

	js.Global().Set("add", js.NewCallback(add))
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Printf("Yes working ...")
	registerCallbacks()

	<-c
}
