package main

import (
	"fmt"
)

func main() {
	// Print an decimal number
	fmt.Println(32)

	// Print binary number
	s := fmt.Sprintf("%d - %b\n", 32, 32)
	fmt.Printf(s)

	// Print a hex number
	fmt.Printf("%d - %x - %#x ", 41, 41, 41)
}
