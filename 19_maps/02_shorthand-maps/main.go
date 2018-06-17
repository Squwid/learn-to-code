package main

import (
	"fmt"
)

func main() {
	// use of composite literal
	var myMap = map[string]int{} // Shorthand way to create a map
	myMap["Ben"] = 20
	fmt.Println(myMap)

	myGreeting := map[string]string{
		"English": "Good morning!",
		"French":  "Bonjour",
	}
	myGreeting["Chinese"] = "Ni Hao!"
	fmt.Println(myGreeting)
}
