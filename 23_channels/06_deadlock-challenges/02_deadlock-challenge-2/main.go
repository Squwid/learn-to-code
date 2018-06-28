package main

import (
	"fmt"
)

func main() {
	// Starting problem, the following code only prints 0 when
	// running so how do you fix that?
	/*	c := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
		fmt.Println(<-c)
	*/

	// This happens because fmt.Println will only recieve the first value
	// and then will quit because it reaches the end of the code block

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // have to close channel when done
	}()

	// Range to not have the program quit after the channel is done
	for d := range c {
		fmt.Println(d)
	}
}
