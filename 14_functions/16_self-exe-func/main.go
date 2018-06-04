package main

import (
	"fmt"
)

// Example of functional programming
func main() {
	// Anonomous self executing function
	//	no reciever no identifier no parameters no return\
	//  params at the end execute the function
	func() {
		fmt.Println("Im driving!")
	}()
}
