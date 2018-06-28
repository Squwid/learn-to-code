package main

import (
	"fmt"
)

/*
	SEMAPHORE EXAMPLE
	Done is a channel of bools and the last go func
		takes both of the bools and discards them, once two go in close(c) gets called ending the program
	Unbuffered channels can only take values in if the channel is ready to receive it
*/

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
