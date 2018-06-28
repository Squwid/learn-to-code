package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Set the max number of cores you are able to use

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // set the max number of cores to use
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait() // Waits until waitgroup goes down to 0. Starts at 2 and each done subtracts one
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("FOO:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()

}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("BAR:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}
