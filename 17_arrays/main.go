package main

import (
	"fmt"
)

// arrays are underlying datastructures that arent generally used

func main() {
	var x [66]string // array
	var x1 []string  // slice

	fmt.Println(len(x))  // 66
	fmt.Println(len(x1)) // 0

	fmt.Printf("%T\n", x)  // [66]string
	fmt.Printf("%T\n", x1) // []string

	// ascii conversion from int to string
	var letters [58]string // array of letters
	for i := 65; i <= 122; i++ {
		letters[i-65] = string(i)
	}

	fmt.Println("\n", letters)
}
