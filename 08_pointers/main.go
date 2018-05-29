package main

import (
	"fmt"
)

func main() {
	referencing()

	x := 5
	zero(&x)
	fmt.Println(x) // x changes

}

func referencing() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a // var b *int = &a

	fmt.Println(b) // Some memory pointer

	// Dereferencing
	fmt.Println(*b) // 43

	// by changing the address of b here it changes the value of a
	*b = 42 // change b to 42
	fmt.Println(a)
}

func zero(z *int) {
	*z = 0
}
