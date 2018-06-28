package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	d := <-c
	fmt.Println("D:", d)

	for n := range c {
		fmt.Println(n)
	}
}
