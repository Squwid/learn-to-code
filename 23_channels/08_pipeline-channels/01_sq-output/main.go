package main

import (
	"fmt"
)

func main() {
	// Set up the pipeline
	c := gen(2, 3)
	out := sq(c)

	// consume the output
	fmt.Println(<-out) // 3
	fmt.Println(<-out) // 9

	//	Another way to do it the same way
	/*
		for n := range sq(gen(2,3)){
			fmt.Println(n)
		}
	*/
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		// taking off of the in chanel by ranging over it
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
