package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("World")
}

func main() {
	defer hello() // moves this statement to the end of closure
	world()
}