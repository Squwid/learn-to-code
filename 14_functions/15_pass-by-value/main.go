package main

import (
	"fmt"
)

func main() {
	age := 44
	fmt.Println(&age) // some address

	changeMe(&age)
	fmt.Println(&age) // some address
	fmt.Println(age)  // 24

	// Second Example
	fmt.Println("\nPart 2:")

	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeMe2(m)
	fmt.Println(m) // [Ben]
}

func changeMe(z *int) {
	fmt.Println(z)  // some address
	fmt.Println(*z) // 44
	*z = 25
	fmt.Println(z)  // some address
	fmt.Println(*z) //25
}

func changeMe2(z []string) {
	z[0] = "Ben"
}
