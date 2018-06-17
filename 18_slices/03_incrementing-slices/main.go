package main

import (
	"fmt"
)

// using a slice and incrementing mySlice[0] by one because it is a refernce type

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0]) // 0
	mySlice[0] = 7
	fmt.Println(mySlice[0]) // 7
	mySlice[0]++
	fmt.Println(mySlice[0]) // 8
}
