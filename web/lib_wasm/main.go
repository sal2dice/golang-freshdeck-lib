package main

import (
	_ "crypto/sha512"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("adder", js.FuncOf(adder))
	<-done
}

func adder(x int, y int) int {
	return x + y
}

/*
func hash(this js.Value, args []js.Value) interface{} {
	h := crypto.SHA512.New()
	h.Write([]byte(args[0].String()))

	return hex.EncodeToString(h.Sum(nil))
}
*/
