package main

import (
	"fmt"
)

/*
	Write a function which takes an integer. The function will have two
	returns. The first return should be the argument divided by 2. The
	second return should be a bool that let's us know wheter or not the
	argument was even
*/

func main() {
	fmt.Println(one(3))
}

func one(inp int) (int, bool) {
	ans := inp / 2
	rem := inp % 2
	var even bool
	if rem == 0 {
		even = true
	}
	return ans, even
}
