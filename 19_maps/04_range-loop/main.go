package main

import (
	"fmt"
)

// iterating over a map

func main() {
	myMap := map[int]string{
		0: "Ben",
		1: "Tom",
		2: "Gid",
		3: "Ted",
	}
	for key, val := range myMap {
		fmt.Printf("Key: %v\tVal: %s\n", key, val)
	}
}
