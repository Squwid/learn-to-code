package main

import (
	"fmt"
	"time"
)

/*
	Simple channels
*/

func main() {
	// c:= make(chan int, 10)  // makes a buffered channels
	c := make(chan int) // unbuffered channel

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // once this is called then the comment below is calle
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // "give me the next value out of the channel"
		}
	}()

	time.Sleep(time.Second)
}
