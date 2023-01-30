package main

import "fmt"

func main() {
	// signed bit intergers, defaultly it will select 32bit min, or 64 or more depending on system
	// n := 42
	// fmt.Printf("default n: %v , %T\n", n, n)

	// var n2 uint16 = 42
	// fmt.Printf("unint16 n2: %v , %T\n", n2, n2)

	//-- arthmetic operations:
	a := 10 // 1010
	b := 3  // 0011
	fmt.Println("adding: ", a+b)
	fmt.Println("subracting: ", a-b)
	fmt.Println("multiply", a*b)
	fmt.Println("divide: ", a/b)
	fmt.Println("reminder", a%b)
	fmt.Println("-------------")

	// > have to same type on performing arthmetic (e.g., can't int and int8)
	// > convert and calculation possible:
	var i1 int = 1
	var i2 int8 = 2
	fmt.Println("converting same ints and adding:", i1+int(i2))
	fmt.Println("-------------")

	// -- operators: and, or, exculsive or and exculsive not
	fmt.Println("and: ", a&b)          // 0010
	fmt.Println("or: ", a|b)           // 1011
	fmt.Println("exculsive or", a^b)   // 1001
	fmt.Println("exculsive not", a&^b) // 0100
	fmt.Println("-------------")

	// -- Bit Shifting:
	a = 8 // 2^3
	fmt.Println("a = 8 or 2^3")
	fmt.Println("bit shifts a left by 3 places: ", a<<3)  // 2^3 * 2^3
	fmt.Println("bit shifts a right by 3 places: ", a>>3) // 2^3 / 2^3
}
