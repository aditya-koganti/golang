package main

import (
	"fmt"
	"strconv"
)

var gi int = 11
var GI int = 12

var (
	actorName         string = "Nikki"
	performanceRating string = "extraordinary"
	personality       string = "good, kind and humorous :)"
	rating            int    = 5
)

func main() {
	var i int = 3
	// most used
	j := 4
	var k int
	fmt.Println("Different types of declarations:")
	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Printf("or can also be wrriten i: value= %v and type = %T", i, i)
	fmt.Println()

	// can be defined again since it is global defined
	var gi int = 1

	// can't be defined again because it was defined CAPITALIZED
	// var GI int = 3

	fmt.Println("getting global value gi: ", gi)

	var f1 = 33.22
	fmt.Printf("golang can also guess type, f1 is %T", f1)
	fmt.Println()
	// special attention needed while converting variable types in golang
	var s1 string = strconv.Itoa(i)
	fmt.Printf("i: %v - converted to string", s1)
	fmt.Println()

}
