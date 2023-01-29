package main

import "fmt"

func main() {
	var b1 bool = false
	fmt.Printf("%v, %T\n", b1, b1)
	// Output: false, bool

	b2 := 1 == 1
	b3 := 1 == 2
	fmt.Printf("%v, %T\n", b2, b2)
	// O: true, bool
	fmt.Printf("%v, %T\n", b3, b3)
	// O: true, bool

	// default value is false
	var b4 bool
	fmt.Printf("%v, %T\n", b4, b4)
	// O: false, bool

}
