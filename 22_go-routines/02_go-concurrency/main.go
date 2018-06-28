package main

import (
	"fmt"
)

// in this case there are three threads running
// 1. main()
// 2. go foo()
// 3. go bar()
// since all of these are running at the same time
// the concurrency causes main to conclude

func main() {
	go foo()
	bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("FOO:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("BAR:", i)
	}
}
