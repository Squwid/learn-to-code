package main

import (
	"errors"
	"log"
)

// ErrBen is an error that is used for testing
var ErrBen = errors.New("this is an idiomatic error")

// example of an idiomatic to make and handle errors
// you want to create errors at the top of the page so they are easy to use
// and you can easily see what the error actually is

func main() {
	_, err := sqrt(-1)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrBen
	}

	return 42, nil
}
