package main

import (
	"fmt"
)

/*
	Example of channels
*/

func main() {
	c := make(chan int)

	// anon func using concurrency
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // when a channel is closed you cannot add values but you can recieve them
	}()

	for n := range c {
		fmt.Println(n)
	}

}
