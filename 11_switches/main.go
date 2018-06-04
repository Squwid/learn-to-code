package main

import (
	"fmt"
)

func main() {
	// Simple switch
	switch "Ben" {
	case "Phil", "Jam":
		fmt.Println("Phil or Jam")
	case "Ian":
		fmt.Println("Ian")
		fallthrough // goes into ben and does not return
	case "Ben":
		fmt.Println("Ben")
	default:
		fmt.Println("Cannot find")
	}

	// Switch w/ No exception
	h := 0
	switch {
	case h == 0:
		fmt.Println('0')
	case h == 1:
		fmt.Println('1')
	}

}

// SwitchOnType gets the type of the object
func SwitchOnType(x interface{}) {
	switch x.(type) { // Called an assert
	case int:
		fmt.Println("int")
	default:
		fmt.Println("Default")
	}
}
