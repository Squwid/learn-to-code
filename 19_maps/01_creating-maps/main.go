package main

import (
	"fmt"
)

// maps are reference types

func main() {
	m := make(map[string]int)
	m["k"] = 0
	m["j"] = 20

	fmt.Println(m) // map[k:1 j:20]

	// This is a problem because k exists as 0 while d does not exist and still
	//	returns to 0
	k := m["k"]
	fmt.Println(k) // 0
	d := m["d"]
	fmt.Println(d) // 0

	// How to fix?
	_, kHere := m["k"]
	_, dHere := m["d"]
	fmt.Println(kHere) // true
	fmt.Println(dHere) // false

	// removing an element from the map
	delete(m, "k") // deletes k from map
	fmt.Println(m) // map[j:20]
}
