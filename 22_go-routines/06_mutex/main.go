package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// example of a data race
// value of counter should be 40 in a perfect world but there
// is a concurrency race making the number be a lot smaller

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		// locks all other threads out except one that is running this
		// good practice to keep braces to show exactly what is being locked
		{
			mutex.Lock()
			counter++
			fmt.Println(s, i, "Counter:", counter)
			// unlocks other threads
			mutex.Unlock()
		}
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
