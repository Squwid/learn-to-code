package main

import (
	"fmt"
)

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	foo := 'a'
	fmt.Println(foo)              // 97
	fmt.Printf("%T\n", foo)       // int32
	fmt.Printf("%T\n", rune(foo)) // int32

	// STRINGS
	a := "This is the first test"

	b := `This""" is the 
	second
	
	
	test`

	c := 'c'
	fmt.Println(a)
	fmt.Println(b)        // Retains spacing
	fmt.Printf("%T\n", c) // int32
}
