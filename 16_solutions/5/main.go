package main

import (
	"fmt"
)

/*
	Write a function, foo, which can be called in each way in the main function
	function called in many ways
*/

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	a := []int{1, 2, 3, 4}
	foo(a...)
	foo()
}

func foo(sf ...int) {
	for _, val := range sf {
		fmt.Print(val, ", ")
	}
}
