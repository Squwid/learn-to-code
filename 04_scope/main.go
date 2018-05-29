package main

import (
	"fmt"
)

// Package scope
var x = 42

func main() {
	// Block scope
	c := 10

	fmt.Println(x)
	fmt.Println(c)

	fmt.Println(max(9))

	scopes()
}

func max(x int) int {
	return 42 + x
}

func scopes() {
	e := 42
	fmt.Println(e)
	{
		fmt.Println(x)
		y := "random string"
		fmt.Println(y)
	}
	// fmt.Println(y) // does not work because out of y's scope
}
