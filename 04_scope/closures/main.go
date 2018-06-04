package main

import (
	"fmt"
)

var x = 0

func main() {
	// P1
	fmt.Println(increment()) // gets an int and not a function
	fmt.Println(increment())

	// P2
	g := 0
	// Anonomous Function
	// Function Expression : Assigning a function to a variable
	increment2 := func() int {
		g++
		return g
	}
	fmt.Println(increment2())
	fmt.Println(increment2())

	// P3
	// Closures
	increment3 := wrapper()
	fmt.Println(increment3()) // 1
	fmt.Println(increment3()) // 2

	d := void()
	d()
}

func void() func() {
	return func() {
		fmt.Println("Here")
	}
}

// P2
func increment() int {
	x++
	return x
}

// P3
// t does not reset in this case
func wrapper() func() int {
	t := 0
	return func() int {
		t++
		return t
	}
}

func test() {
	var v = 10
	{
		fmt.Println(v)
	}
}
