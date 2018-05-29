package main

import (
	"fmt"
)

const (
	pi       = 3.14
	language = "Go"
)

// IOTAS
const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

// Bitwise example
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)  -- 1 << 10 = 1024
	MB = 1 << (iota * 10) // 1 << (2 * 10)  -- 1 << 20 = 1,048,576
)

const p = "constant string"

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println(pi)
	fmt.Println(language)

	fmt.Printf("KB: %v\t MB:%v\n", KB, MB)
}

// a CONSTANT is a simple unchanging value
