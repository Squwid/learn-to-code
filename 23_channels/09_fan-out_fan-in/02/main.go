package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// FAN OUT
	// distribute the sq work across two goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}

}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	fmt.Println("Type: %T", cs) // fyi
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c) // exaple of passing an argument into an anon func
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
