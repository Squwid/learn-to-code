package main

import (
	"fmt"
)

func main() {
	// This code below does not work because of deadlock,
	// the way to fix it is even further below
	/*
		c := make(chan int)
		c <- 1
		fmt.Println(<-c)
	*/

	// There needs to be something listening to the channel in order to recieve
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
