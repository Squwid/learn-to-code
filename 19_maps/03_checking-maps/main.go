package main

import (
	"fmt"
)

func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos dias",
		3: "Bongiorno",
	}
	fmt.Println(myGreeting)

	// example of a shortened if statement again
	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2) // delete key 2 from myGreeting
		fmt.Println("Val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("Value does not exist")
	}
}
