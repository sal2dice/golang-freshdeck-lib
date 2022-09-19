package lib_js

// This exports an add function.
// It takes in two 32-bit integer values
// And returns a 32-bit integer value.
// To make this function callable from JavaScript,
// we need to add the: "export add" comment above the function
//export add
func adder(x int, y int) int {
	return x + y
}
