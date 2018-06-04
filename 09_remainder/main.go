package main

import "fmt"

func main() {
	x := 13 % 2 // remainder
	fmt.Println(x)

	// see if a number is odd or even
	num := 9
	if num%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
