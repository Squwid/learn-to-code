package main

import (
	"fmt"
)

func main() {
	n := average(32, 34, 35, 66, 99)
	fmt.Println(n)

	// Variadic arguments
	data := []float64{43, 33, 23, 56, 76}
	_ = average(data...) // sends the entire list

}

// Notice where the ... are

// Variadic function
func average(sf ...float64) float64 {
	fmt.Println(sf)         // [32, 35, ...]
	fmt.Printf("%T \n", sf) // []float64
	for _, val := range sf {
		fmt.Println(val)
	}
	return 0.0
}
