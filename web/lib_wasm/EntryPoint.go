package main

import (
	_ "crypto/sha512"
	"syscall/js"
)

// from https://binx.io/2022/04/22/golang-webassembly/
//  GOOS=js GOARCH=wasm go build -o html/lib_freshdeck.wasm lib_wasm/EntryPoint.go

//cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./html/

type UsefulObject struct {
	a  string
	bc string
}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("adder", js.FuncOf(adder))
	<-done
}

func adder(x int, y int) int {
	return x + y
}
