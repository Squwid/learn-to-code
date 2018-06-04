package main

import (
	"fmt"
)

/*
	Write a function with one variadic parameter that finds the greatest
	number in a list of numbers
*/

func main() {
	n := max(1, 2, 3, 4, 5, 6, 7, 8, 9, 20)
	fmt.Println(n)
}

// Variadic parameter for max
func max(sf ...int64) int64 {
	var max int64
	for _, val := range sf {
		if val > max {
			max = val
		}
	}
	return max
}
