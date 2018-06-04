package main

import "fmt"

func main() {

	// for init; cond; post {}
	for i := 0; i <= 5; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		}
		//fmt.Println(i)
	}

	// while loop said  (for k is less than 10)
	k := 0
	for k < 10 {
		fmt.Println(k)
		k++
	}
}
