package main

import (
	"fmt"
	"sync"
)

/*
	This does not work because you cannot add to a variable in two different
	go routines, see the 02_wait-group for the correct way to do it
*/

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
