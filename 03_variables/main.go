package main

import "fmt"

// Declaration
var z string

func main() {
	// Shorthand
	a := 10
	b := "words as string"

	fmt.Printf("%v\n", a) // prints 10
	fmt.Printf("%v\n", b) // prints words as string

	fmt.Printf("%T\n", a) // prints int
	fmt.Printf("%T\n", b) // prints string
}
