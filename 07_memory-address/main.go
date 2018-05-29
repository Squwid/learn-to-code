package main

import "fmt"

func main() {
	a := 42

	fmt.Println("a - ", a)
	fmt.Println("a's memory address-", &a)
	fmt.Printf("%d\n", &a)

	// Using memory addresses
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	fmt.Println("Meters ", meters)
}
