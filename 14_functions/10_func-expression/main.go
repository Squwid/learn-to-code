package main

import "fmt"

func main() {

	// example of anon function
	greeting := func() {
		fmt.Println("Hello world!")
	}
	greeting()

	// func expressions
	greeting2 := makeGreeter()
	fmt.Printf("%T\n", greeting2())
}

func makeGreeter() func() string {
	fmt.Println("Before return") // Gets called as greeting 2 is created
	return func() string {
		return "Hello world"
	}
}
