package main

import "fmt"

func one() {
	fmt.Println("Hello World")
}

func two() {
	name := "Ben"
	fmt.Printf("Hello %s\n", name)
}

// Gets the name from
func three() {
	var name string
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello %s\n", name)
}

func four(one int, two int) {
	fmt.Println(one % two)
}

func main() {

}
